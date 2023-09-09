package main

import (
	"github.com/solace-labs/skeyn/node"
	"github.com/solace-labs/skeyn/utils"
)

func main() {
	node := node.NewNode()
	// TODO: Pass context from here

	node.Start(utils.GetCliConfig())
	select {}
}
