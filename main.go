package main

import "libp2p-playground/node"

func main() {
	node := node.NewNode()
	// TODO: Pass context from here
	node.Start()
	select {}
}
