package node

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	smartcontract "github.com/solace-labs/skeyn/smart_contract"
	"github.com/solace-labs/skeyn/squad"
	"github.com/solace-labs/skeyn/utils"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
)

const (
	RENDEVOUS_STRING = "SOLACE_PROTOCOL"
)

const networkCidString = "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"

var (
	buf, _      = hex.DecodeString(networkCidString)
	mHashBuf, _ = multihash.EncodeName(buf, "sha1")
	mHash, _    = multihash.FromHexString(hex.EncodeToString(mHashBuf))
	networkCid  = cid.NewCidV1(cid.Raw, mHash)
)

type SquadM map[common.WalletAddress]*squad.Squad

type Node struct {
	host *host.Host
	kdht *dht.IpfsDHT

	squad SquadM

	smartContract smartcontract.NetworkState

	proto.UnimplementedTransactionServiceServer
}

func NewNode() *Node {
	return &Node{}
}

func (n Node) h() host.Host {
	return *n.host
}

func (n Node) ShortID() string {
	return n.h().ID().Pretty()[:5]
}

func (n Node) PeerID() peer.ID {
	return n.h().ID()
}

func (n *Node) Start(ctx context.Context, config utils.Config) {

	pub, _ := config.GetPrivKey().GetPublic().Raw()
	log.Println("Public Key", hex.EncodeToString(pub))

	n.squad = make(SquadM)

	n.CreateHost(ctx, config.GetPort(), config.GetPrivKey())

	n.CreateDHT(ctx, dht.Mode(dht.ModeServer)) // Always run in server mode for peer discovery

	n.ConnectBootstrapPeers(ctx, config.GetPeers()) // Connect to bootstrap peers, who run DHT in Server (Full) mode

	go n.discoverProviders(ctx) // Keep discovering new providers who might offer the same

	err := n.kdht.Bootstrap(ctx) // Important to bootstrap after finding other providers
	if err != nil {
		panic(err)
	}

	// go n.FindPeers(ctx) // Not required as long as you have one bootstrap node

	n.SetupNotifications()

	n.smartContract = &smartcontract.TestContract{}

	n.setupMessageRecieverHandler()
	// Instead of having a channel, you could just call the function on the target squad
	// TODO: Rethink this 'incoming channel'  approach

	// TODO: Setup an updater which keeps checking if the network state is valid with the node

	if config.GetShouldRunExternalRPCServer() {
		n.SetupGRPC(ctx)
	}

}

func (n *Node) CreateHost(ctx context.Context, port int, privKey crypto.PrivKey) {
	listen, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	host, err := libp2p.New(libp2p.ListenAddrs(listen), libp2p.Identity(privKey))
	if err != nil {
		panic(err)
	}
	n.host = &host
	// Log the hosting addresses
	for _, addr := range n.h().Addrs() {
		log.Printf("%s/p2p/%s\n", addr, n.h().ID())
	}
	log.Println("Host Created: ", host.ID().Pretty(), "port: ", port)
}

func (n *Node) CreateDHT(ctx context.Context, options ...dht.Option) {
	kdht, err := dht.New(ctx, n.h(), options...)
	if err != nil {
		panic(err)
	}
	n.kdht = kdht
	err = kdht.Provide(ctx, networkCid, true)
	if err != nil {
		log.Println("[WARN] Providing DHT: ", err)
	}

}

func (n *Node) ConnectBootstrapPeers(ctx context.Context, addrs utils.AddrList) {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		peerInfo, _ := peer.AddrInfoFromP2pAddr(addr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := n.h().Connect(ctx, *peerInfo); err != nil {
				log.Println("ERER", err)
			} else {
				log.Println("Connected to a bootstrap peer")
			}
		}()
	}
}

func (n *Node) FindPeers(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			go n.GetRandomPeers()
		}
	}
}

func (n *Node) discoverProviders(ctx context.Context) {
	ch := n.kdht.FindProvidersAsync(ctx, networkCid, 0)
	for {
		// Run every 2 seconds
		time.Sleep(2 * time.Second)
		// log.Println(n.h().Network().Peers())
		select {
		case <-ctx.Done():
			return
		case peer := <-ch:
			if peer.ID == "" {
				continue
			}
			if peer.ID == (*n.host).ID() {
				continue
			}
			// TODO: Make this more strict. Connect only if certain criteria matches
			err := n.h().Connect(ctx, peer)
			if err != nil {
				log.Println("Error connecing to peer", peer.ID)
				continue
			}
			log.Println((*n.host).ID().String(), " [PROVIDER] ", peer.ID)
			_ = n.kdht.Bootstrap(ctx)
		}
	}
}

func (n *Node) GetRandomPeers() {
	randomKey := make([]byte, 32)
	_, err := rand.Read(networkCid.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	peers, err := n.kdht.GetClosestPeers(ctx, string(randomKey))
	if err != nil {
		log.Println(err)
	}
	for _, peer := range peers {
		if peer == "" {
			continue
		}
		if peer == n.h().ID() {
			continue
		}
		if n.h().Network().Connectedness(peer) == network.Connected {
			continue
		}
		log.Println("Found peer", peer)
		_, err := n.h().Network().DialPeer(ctx, peer)
		if err != nil {
			log.Println("Error connecting to peer")
			panic(err)
		}
		log.Println("Connected to ", peer)
	}

}

func (n *Node) SetupNotifications() {
	n.h().Network().Notify(&network.NotifyBundle{
		ConnectedF: func(x network.Network, c network.Conn) {
			log.Println(n.h().ID().Pretty(), "{ CONNECTED }", c.RemotePeer().Pretty())
		},
		DisconnectedF: func(x network.Network, c network.Conn) {
			log.Println(n.h().ID().Pretty(), "{ DISCONNECTED }", c.RemotePeer().Pretty())
			log.Println("CONNECTED TO:", len(n.h().Network().Peers()))
		},
	})
}
