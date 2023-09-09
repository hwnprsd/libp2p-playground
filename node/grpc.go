package node

import (
	"context"
	"log"
	"net"
	"net/http"

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
