package node

import (
	"context"
	"github.com/solace-labs/skeyn/common"
	proto "github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/utils"
	"log"

	"github.com/libp2p/go-libp2p/core/network"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// 1. Check if the wallet is under management
// 2. Check if the signer of the message is authorzed
// 3. Check if the type of message is valid (Start DKG vs Sign)
// 4. Signal the squad to do what's required
// Called by external agents
func (n *Node) SendTransaction(ctx context.Context, req *proto.Transaction) (*proto.TransactionResponse, error) {
	isInvalidRequest := req.Type == "" ||
		req.Payload == nil ||
		req.Payload.WalletAddress == nil ||
		req.Payload.Signature == nil ||
		req.Payload.Data == nil

	if isInvalidRequest {
		return &proto.TransactionResponse{Success: false, Msg: "Invalid Request"}, nil
	}

	sig := req.Payload.Signature

	signDataHash := ethcrypto.Keccak256Hash(req.Payload.Data)
	pubKey, err := ethcrypto.SigToPub(signDataHash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}

	address := ethcrypto.PubkeyToAddress(*pubKey)
	_ = address

	// Check if address and wallet address are a part of the squad

	// TODO: Broadcast random shit to peers
	log.Println("TYPE", req.Type)
	if req.Type == "1" {
		n.squad.InitKeygen(ctx)
	} else {
		n.squad.InitSigning(ctx, []byte("YEET"))
	}

	// Your logic here
	return &proto.TransactionResponse{Success: true, Msg: "ok"}, nil
}

// Message sink
func (n *Node) setupMessageRecieverHandler(ctx context.Context) <-chan common.IncomingMessage {
	ch := make(chan common.IncomingMessage)
	go func() {
		// Setup recievers
		n.h().SetStreamHandler(common.DKG_PROTOCOL, func(s network.Stream) {
			data, err := utils.ReadStream(s)
			if err != nil {
				log.Println("Error Reading Stream Data")
				log.Println(err)
			}

			nodeMessage := common.NodeMessage{
				PeerID:   s.Conn().RemotePeer(),
				Data:     data,
				Protocol: common.DKG_PROTOCOL,
			}
			ch <- nodeMessage
		})

		n.h().SetStreamHandler(common.SIGNING_PROTOCOL, func(s network.Stream) {
			data, err := utils.ReadStream(s)
			if err != nil {
				log.Println("Error Reading Stream Data")
				log.Println(err)
			}

			nodeMessage := common.NodeMessage{
				PeerID:   s.Conn().RemotePeer(),
				Data:     data,
				Protocol: common.SIGNING_PROTOCOL,
			}
			ch <- nodeMessage
		})

	}()
	return ch
}

// Message Export Hub for the node
func (n *Node) setupOutgoingMessageHandler(ctx context.Context) chan common.OutgoingMessage {
	ch := make(chan common.OutgoingMessage)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-ch:
				stream, err := n.h().NewStream(ctx, msg.GetPeerID(), msg.GetProtocol())
				if err != nil {
					log.Println("Error creating stream", err)
					continue
				}

				if err := utils.WriteStream(stream, msg.GetData()); err != nil {
					log.Println("Error writing to stream", err)
					continue
				}
			}
		}
	}()
	return ch
}
