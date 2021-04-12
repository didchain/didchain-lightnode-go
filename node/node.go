package node

import (
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/filter"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/didchain/didchain-lightnode-go/config"
	"github.com/didchain/didchain-lightnode-go/loginUam"
	"github.com/didchain/didchain-lightnode-go/user/storage"
	"log"
	"sync"
)

type LightNode struct {
	conf   *config.NodeConfig
	worker *Worker
}

var LightNodeInstance *LightNode
var oncedo sync.Once

func NodeInstance(cfg *config.NodeConfig) *LightNode {
	oncedo.Do(func() {
		LightNodeInstance = NewNode(cfg)
	})

	return LightNodeInstance
}

func NewNode(cfg *config.NodeConfig) *LightNode {

	opts := opt.Options{
		Strict:      opt.DefaultStrict,
		Compression: opt.NoCompression,
		Filter:      filter.NewBloomFilter(10),
	}

	db, err := leveldb.OpenFile(cfg.DatabasePath, &opts)
	if err != nil {
		panic(err)
	}

	node := &LightNode{
		conf: cfg,
		worker: &Worker{port: cfg.ListenPort, loginwebport: cfg.LoginPort, storage: storage.NewStorage(db),
			admin: config.LoadAdminUser(cfg.AdminUserDb), cfg: cfg, sessionStorage: loginUam.NewSessStorage(cfg)},
	}

	return node
}

func (sn *LightNode)GetUserStorage() *storage.Storage  {
	return sn.worker.storage
}

func (sn *LightNode) Start() {
	log.Println("begin to start web")
	sn.worker.StartWebDaemon()

	log.Println("begin to start login web")
	sn.worker.StartLoginWebDaemon()
}

func (sn *LightNode) Stop() {
	sn.worker.StopWorker()
}
