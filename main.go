package main

func main() {
	node := NewNode()
	// TODO: Pass context from here
	node.Start()
	select {}
}
