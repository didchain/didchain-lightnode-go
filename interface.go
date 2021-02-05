package main

type NodeConfig struct {
	ListenPort int `json:"listen_port"`
}

type LightNode struct {
}

func NewNode(cfg NodeConfig) *LightNode {
	node := &LightNode{}
	return node
}

func (sn *LightNode) Start() {
	//http server
	//only bind localhost: for safety
	//new go thread for each request
}
