package config

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kprc/nbsnetwork/tools"
	"log"
	"os"
	"path"
	"sync"
)

const lightNodeConfFile = "did-conf-file.json"
const leveldbDir = "db"

type NodeConfig struct {
	ListenPort     int    `json:"listen_port"`
	LoginPort      int    `json:"login_port"`
	DatabasePath   string `json:"database_path"`
	AdminUserDb    string `json:"admin_user_db"`
	LoginUrl       string `json:"login_url"`
	SessiontimeOut int64  `json:"sessiontime_out"`
	RedirUrl       string `json:"redir_url"`
}

func LightNodeHome() string {
	home, _ := tools.Home()
	return path.Join(home, ".didHome")
}

func LightNodeConfFile() string {
	return path.Join(LightNodeHome(), lightNodeConfFile)
}

func InitNodeConf() *NodeConfig {

	var cfg *NodeConfig

	cfgfile := LightNodeConfFile()
	if tools.FileExists(cfgfile) {
		if data, err := tools.OpenAndReadAll(cfgfile); err != nil {
			panic(err.Error())
		} else {
			cfg = &NodeConfig{}
			if err = json.Unmarshal(data, cfg); err != nil {
				panic(err.Error())
			}
		}

	} else {
		cfg = &NodeConfig{
			ListenPort:     60999,
			LoginPort:      60998,
			DatabasePath:   path.Join(LightNodeHome(), leveldbDir),
			AdminUserDb:    path.Join(LightNodeHome(), "adminUser.db"),
			LoginUrl:       "http://39.99.198.143:60998/api/verify",
			RedirUrl:       "http://39.99.198.143:60998/index.html",
			SessiontimeOut: 1200000,
		}

		cfg.Save()
	}

	return cfg
}

func (cfg *NodeConfig) Save() {
	j, _ := json.MarshalIndent(*cfg, " ", "\t")

	if !tools.FileExists(LightNodeHome()) {
		os.MkdirAll(LightNodeHome(), 0755)
	}

	if err := tools.Save2File(j, LightNodeConfFile()); err != nil {
		log.Println("save to ", LightNodeConfFile(), " failed")
	}
}

type AdminUser struct {
	Lock     sync.Mutex
	savePath string
	EthAddr  []string `json:"eth_addr"`
}

var GAdminUser *AdminUser

func LoadAdminUser(adminfile string) *AdminUser {

	au := &AdminUser{savePath: adminfile}
	if data, err := tools.OpenAndReadAll(adminfile); err != nil {
		GAdminUser = au
		return au
	} else {
		var addrs []string
		if err = json.Unmarshal(data, &addrs); err != nil {
			panic(err.Error())
		}

		for i := 0; i < len(addrs); i++ {
			au.EthAddr = append(au.EthAddr, addrs[i])
		}

		GAdminUser = au

		return au
	}
}

func (au *AdminUser) Save() {
	addrs := au.listUser()

	data, _ := json.MarshalIndent(addrs, " ", "\t")
	tools.Save2File(data, au.savePath)
}

func (au *AdminUser) AddUser(id string) error {
	if b := common.IsHexAddress(id); !b {
		return errors.New("not a correct eth address")
	}

	au.Lock.Lock()
	defer au.Lock.Unlock()

	for i := 0; i < len(au.EthAddr); i++ {
		if au.EthAddr[i] == id {
			return errors.New("duplicate user")
		}
	}

	au.EthAddr = append(au.EthAddr, id)

	au.Save()

	return nil
}

func (au *AdminUser) DelUser(id string) error {
	if b := common.IsHexAddress(id); !b {
		return errors.New("not a correct eth address")
	}
	au.Lock.Lock()
	defer au.Lock.Unlock()

	fordel := -1

	for i := 0; i < len(au.EthAddr); i++ {
		if au.EthAddr[i] == id {
			fordel = i
		}
	}

	if fordel == -1 {
		return errors.New("no address in db")
	}
	l := len(au.EthAddr)
	au.EthAddr[fordel] = au.EthAddr[l-1]

	au.EthAddr = au.EthAddr[:l-1]

	au.Save()

	return nil
}

func (au *AdminUser) ListUser() []string {
	au.Lock.Lock()
	defer au.Lock.Unlock()

	var r []string

	for i := 0; i < len(au.EthAddr); i++ {
		r = append(r, au.EthAddr[i])
	}

	return r
}
func (au *AdminUser) listUser() []string {

	var r []string

	for i := 0; i < len(au.EthAddr); i++ {
		r = append(r, au.EthAddr[i])
	}

	return r
}
