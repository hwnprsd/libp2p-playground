package main

import (
	"libp2p-playground/node"
	"libp2p-playground/utils"
)

func main() {
	node := node.NewNode()
	// TODO: Pass context from here

	node.Start(utils.GetCliConfig())
	select {}
}
