package node

import (
	"context"
	"encoding/hex"
	"log"

	"github.com/solace-labs/skeyn/common"
	proto "github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/utils"

	"github.com/libp2p/go-libp2p/core/network"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// 1. Check if the wallet is under management
// 2. Check if the signer of the message is authorzed
// 3. Check if the type of message is valid (Start DKG vs Sign)
// 4. Signal the squad to do what's required
// Called by external agents
func (n *Node) HandleTransaction(ctx context.Context, req *proto.Transaction) (*proto.TransactionResponse, error) {
	walletAddress := ethcommon.HexToAddress(req.Payload.WalletAddress)
	signature := ethcommon.Hex2Bytes(req.Payload.Signature)
	data := ethcommon.Hex2Bytes(req.Payload.Data)

	isInvalidRequest := req.Type == "" ||
		req.Payload == nil ||
		walletAddress.Bytes() == nil ||
		signature == nil ||
		data == nil

	if isInvalidRequest {
		return &proto.TransactionResponse{Success: false, Msg: "Invalid Request"}, nil
	}

	signDataHash := ethcrypto.Keccak256Hash(data)
	pubKey, err := ethcrypto.SigToPub(signDataHash.Bytes(), signature)
	if err != nil {
		return &proto.TransactionResponse{Success: false, Msg: err.Error()}, nil
	}

	address := ethcrypto.PubkeyToAddress(*pubKey)
	log.Println("Address - ", address.Hex())

	// Check if address and wallet address are a part of the squad

	// TODO: Broadcast random shit to peers
	if req.Type == "1" {
		n.squad.InitKeygen(ctx)
	} else {
		// Verify Incoming Message
		n.squad.InitSigning(ctx, data)
	}

	key := hex.EncodeToString(data)

	// Your logic here
	return &proto.TransactionResponse{Success: true, Msg: key}, nil
}

func (n *Node) HandleSigRetrieval(ctx context.Context, req *proto.SignatureRetrieval) (*proto.TransactionResponse, error) {
	sig, err := n.squad.GetSig([]byte(req.Key))
	if err != nil {

		return &proto.TransactionResponse{Success: false, Msg: "Error fetching squad sig"}, err
	}
	return &proto.TransactionResponse{Success: true, Msg: hex.EncodeToString(sig)}, nil
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
