package node

import (
	"encoding/json"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/filter"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/didchain/didchain-lightnode-go/user/storage"
	"github.com/kprc/nbsnetwork/tools"

	"log"
	"os"
	"path"
)

const lightNodeConfFile = "did-conf-file.json"
const leveldbDir = "db"

type NodeConfig struct {
	ListenPort int `json:"listen_port"`
	DatabasePath string `json:"database_path"`
}

type LightNode struct {
	conf *NodeConfig
	worker *Worker
}

func LightNodeHome() string  {
	home,_:=tools.Home()
	return path.Join(home,".didHome")
}

func LightNodeConfFile()  string {
	return path.Join(LightNodeHome(), lightNodeConfFile)
}


func InitNodeConf() *NodeConfig {

	var cfg *NodeConfig

	cfgfile := LightNodeConfFile()
	if tools.FileExists(cfgfile){
		if data,err:=tools.OpenAndReadAll(cfgfile);err!=nil{
			panic(err.Error())
		}else{
			cfg = &NodeConfig{}
			if err=json.Unmarshal(data,cfg);err!=nil{
				panic(err.Error())
			}
		}

	}else{
		cfg = &NodeConfig{
			ListenPort: 50999,
			DatabasePath: path.Join(LightNodeHome(),leveldbDir),
		}

		cfg.Save()
	}

	return cfg
}

func (cfg *NodeConfig)Save()  {
	j,_:=json.MarshalIndent(*cfg," ","\t")

	if !tools.FileExists(LightNodeHome()){
		os.MkdirAll(LightNodeHome(),0755)
	}

	if err:=tools.Save2File(j, LightNodeConfFile());err!=nil{
		log.Println("save to ", LightNodeConfFile()," failed")
	}
}



func NewNode(cfg *NodeConfig) *LightNode {

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
		conf:cfg,
		worker: &Worker{port: cfg.ListenPort,storage: storage.NewStorage(db)},
	}

	return node
}

func (sn *LightNode) Start() {
	sn.worker.StartWebDaemon()
}

func (sn *LightNode) Stop()  {
	sn.worker.StopWebDaemon()
}
