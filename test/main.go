package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/didchain/didCard-go/account"
	"github.com/didchain/didchain-lightnode-go/protocol"
)

func main()  {
	//testinterface()


	testloadwallet()

}

func testloadwallet()  {
	//w,_:=account.NewWallet("123")

	//fmt.Println("-----",hex.EncodeToString(w.PrivKey()))
	wp:="/Users/rickeyliao/gowork/src/github.com/didchain/didchain-lightnode-go/test/testwallet"
	//w.SaveToPath(wp)

	//fmt.Println(w.String())

	//w.Close()

	wl,err:=account.LoadWallet(wp)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(wl.String())

	err = wl.Open("123")
	if err!=nil{
		panic(err.Error())
	}

	fmt.Println(hex.EncodeToString(wl.PrivKey()))

	fmt.Println("open wallet success")
}

func testverifysig()  {
	w,_:=account.NewWallet("123")
	did := w.Did()

	vr:=&protocol.VerfiyPlainMsg{
		DID: did.String(),
		TimeStamp: 1612312312345,
		Latitude: 111.21,
		Longitude: 21.23,
	}

	sig:=w.SignJson(vr)

	req := &protocol.Request{
		Action: protocol.VerifySignature,
	}

	svr:=&protocol.VerifyReq{
		Signature: base58.Encode(sig),
		VerfiyPlainMsg:*vr,
	}

	req.PayLoad = svr

	j,_:=json.MarshalIndent(req," ","\t")

	fmt.Println(string(j))

	b:=account.VerifySig(did,sig,vr)

	fmt.Println(b)
}


func testinterface()  {
	req:=&protocol.Request{Action: protocol.VerifySignature}



	var sigbytes = make([]byte,64)
	rand.Read(sigbytes)

	vr:=&protocol.VerifyReq{
		Signature: base58.Encode(sigbytes),
		VerfiyPlainMsg:protocol.VerfiyPlainMsg{DID: base58.Encode(sigbytes[32:]),
			TimeStamp: 1234567123456,
			Latitude: 112.22,
			Longitude: 12.22,
		},
	}

	req.PayLoad = vr

	j,_:=json.MarshalIndent(req," ","\t")

	fmt.Println(string(j))

	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.MethodNotCorrect], protocol.MethodNotCorrect).String())

	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.InterIOErr], protocol.InterIOErr).String())
	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.UnmarshalJsonErr], protocol.UnmarshalJsonErr).String())

	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.ActionErr], protocol.ActionErr).String())

	resp:=&protocol.VerifyResp{Signature: vr}

	fmt.Println(protocol.ResponseSuccess(resp).String())


	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.SignatureNotCorrect], protocol.SignatureNotCorrect).String())
}