package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kprc/nbsnetwork/tools"
)

type HomeEntranceSignature struct {

	Signature string `json:"signature"`
	DID string `json:"did"`
}

type HomeEntranceSignatureResult struct {
	Signature *HomeEntranceSignature `json:"signature"`
	ResultCode int `json:"result_code"`
	Result bool `json:"result"`
}


func main() {
	//cfg := NodeConfig{}
	//node := NewNode(cfg)
	//
	//node.Start()

	//TODO wait user signal

	sig:=make([]byte,33)
	rand.Read(sig)

	did:=make([]byte,48)
	rand.Read(did)

	hes:=&HomeEntranceSignature{
		TimeStamp: tools.GetNowMsTime(),
		Latitude: 22.223344,
		Longitude: 44.223322,
		Signature: base64.StdEncoding.EncodeToString(sig),
		DID: base64.StdEncoding.EncodeToString(did),
	}

	hesr:=&HomeEntranceSignatureResult{
		Signature: hes,
		Result: true,
		ResultCode: 0,
	}

	j,_:=json.MarshalIndent(hesr," ","\t")

	fmt.Println(string(j))

}
