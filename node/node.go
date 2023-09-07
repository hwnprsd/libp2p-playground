package node

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"libp2p-playground/proto"
	"libp2p-playground/utils"
	"log"
	"sync"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
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

type Node struct {
	host *host.Host
	kdht *dht.IpfsDHT
	ps   *pubsub.PubSub

	proto.UnimplementedTransactionServiceServer
}

func NewNode() *Node {
	return &Node{}
}

func (n Node) h() host.Host {
	return *n.host
}

func (n Node) shortID() string {
	return n.h().ID().Pretty()[:5]
}

func (n Node) PeerID() peer.ID {
	return n.h().ID()
}

func (n *Node) Start(config utils.Config) {
	ctx := context.Background()

	pub, _ := config.GetPrivKey().GetPublic().Raw()
	log.Println("Public Key", hex.EncodeToString(pub))

	n.CreateHost(ctx, config.GetPort(), config.GetPrivKey())

	n.CreateDHT(ctx, dht.Mode(dht.ModeServer)) // Always run in server mode for peer discovery

	n.ConnectBootstrapPeers(ctx, config.GetPeers()) // Connect to bootstrap peers, who run DHT in Server (Full) mode

	go n.discoverProviders(ctx) // Keep discovering new providers who might offer the same

	n.kdht.Bootstrap(ctx) // Important to bootstrap after finding other providers

	// go n.FindPeers(ctx) // Not required as long as you have one bootstrap node

	n.SetupNotifications()

	n.SetupPubSub(ctx)

	if config.GetShouldRunExternalRPCServer() {
		n.SetupGRPC()
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
		log.Println(fmt.Sprintf("%s/p2p/%s", addr, n.h().ID()))
	}
	log.Println("Host Created: ", host.ID().Pretty(), "port: ", port)
}

func (n *Node) CreateDHT(ctx context.Context, options ...dht.Option) {
	kdht, err := dht.New(ctx, n.h(), options...)
	if err != nil {
		panic(err)
	}
	n.kdht = kdht
	kdht.Provide(ctx, networkCid, true)

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
		// case <-ctx.Done():
		// 	return
		case peer := <-ch:
			if peer.ID == "" {
				continue
			}
			if peer.ID == (*n.host).ID() {
				continue
			}
			// TODO: Make this more strict. Connect only if certain criteria matches
			n.h().Connect(ctx, peer)
			log.Println((*n.host).ID().String(), " [PROVIDER] ", peer.ID)
			n.kdht.Bootstrap(ctx)
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
			x.ClosePeer(c.RemotePeer())
		},
		DisconnectedF: func(x network.Network, c network.Conn) {
			log.Println(n.h().ID().Pretty(), "{ DISCONNECTED }", c.RemotePeer().Pretty())
		},
	})
}
