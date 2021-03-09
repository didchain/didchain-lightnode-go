package node

import (
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/filter"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/didchain/didchain-lightnode-go/config"
	"github.com/didchain/didchain-lightnode-go/user/storage"
)

type LightNode struct {
	conf   *config.NodeConfig
	worker *Worker
}

func NewNode(cfg *config.NodeConfig) *LightNode {

	//if !tools.FileExists(cfg.DatabasePath){
	//	os.MkdirAll(cfg.DatabasePath,0755)
	//}

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
		worker: &Worker{port: cfg.ListenPort, storage: storage.NewStorage(db),
			admin: config.LoadAdminUser(cfg.AdminUserDb)},
	}

	return node
}

func (sn *LightNode) Start() {
	sn.worker.StartWebDaemon()
}

func (sn *LightNode) Stop() {
	sn.worker.StopWebDaemon()
}
