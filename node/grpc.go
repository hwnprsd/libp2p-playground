package node

import (
	"context"
	"encoding/hex"
	"log"
	"net"
	"net/http"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	proto "github.com/solace-labs/skeyn/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: Handle errors
func (n *Node) SetupGRPC() {
	lis, err := net.Listen("tcp", ":5123")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterTransactionServiceServer(grpcServer, n)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx := context.Background()

	if err := proto.RegisterTransactionServiceHandlerFromEndpoint(ctx, mux, "localhost:5123", opts); err != nil {
		panic(err)
	}

	go func() {
		log.Println("Running GRPC/HTTP on port", 5050)
		log.Fatal(http.ListenAndServe(":5050", mux))
	}()

	log.Println("Running GRPC Server on port", 5123)
	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()
}

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

func (n *Node) HandleCreateRule(ctx context.Context, data *proto.CreateRuleData) (*proto.TransactionResponse, error) {
	return nil, nil
}
