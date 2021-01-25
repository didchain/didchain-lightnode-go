package main

func main() {
	cfg := NodeConfig{}
	node := NewNode(cfg)

	node.Start()

	//TODO wait user signal
}
