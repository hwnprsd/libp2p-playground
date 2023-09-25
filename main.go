package main

import (
	"context"

	"github.com/solace-labs/skeyn/node"
	"github.com/solace-labs/skeyn/utils"
)

func main() {
	node := node.NewNode()
	// TODO: Pass context from here

	node.Start(context.TODO(), utils.GetCliConfig())
	select {}
}
