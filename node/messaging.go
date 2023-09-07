package node

import (
	"context"
	proto "libp2p-playground/proto"
)

// Called by external agents
func (n *Node) SendTransaction(ctx context.Context, req *proto.Transaction) (*proto.TransactionResponse, error) {
	// Your logic here
	return &proto.TransactionResponse{Success: true, Msg: "ok"}, nil
}
