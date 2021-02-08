package main

import (
	"context"
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/didchain/didCard-go/account"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"time"
)

type Worker struct {
	port int
	webserver *http.Server
}


func (w *Worker)StartWebDaemon() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ed25519/signatureVerify", ed25519Verify)

	addr := "127.0.0.1:" + strconv.Itoa(w.port)

	log.Println("Web Server Start at", addr)

	w.webserver = &http.Server{Addr: addr, Handler: mux}

	go log.Fatal(w.webserver.ListenAndServe())
}

func (w *Worker)StopWebDaemon() {
	if w.webserver == nil{
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	w.webserver.Shutdown(ctx)

	w.webserver = nil

	log.Println("Web Server Stopped")
}

func ed25519Verify(w http.ResponseWriter,r *http.Request)  {
	if r.Method != "POST"{
		w.WriteHeader(500)
		w.Write(ResponseError(ErrDesc[MethodNotCorrect],MethodNotCorrect).Bytes())
		return
	}

	rbytes,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		w.WriteHeader(200)
		w.Write(ResponseError(ErrDesc[InterIOErr],InterIOErr).Bytes())
		return
	}

	vr:=&VerifyReq{}

	req:=&Request{
		PayLoad: vr,
	}

	if err:=json.Unmarshal(rbytes,req);err!=nil{
		w.WriteHeader(200)
		w.Write(ResponseError(ErrDesc[UnmarshalJsonErr],UnmarshalJsonErr).Bytes())
		return
	}

	if req.Action != VerifySignature{
		w.WriteHeader(200)
		w.Write(ResponseError(ErrDesc[ActionErr],ActionErr).Bytes())
		return
	}

	b:=account.VerifySig(account.ID(vr.DID),base58.Decode(vr.Signature),vr.VerfiyPlainMsg)

	if b{
		w.WriteHeader(200)
		resp:=&VerifyResp{Signature: vr,
			ResultCode: Sucess,
			Result: true}
		w.Write(ResponseSuccess(resp).Bytes())
	}else{
		w.WriteHeader(200)
		resp:=&VerifyResp{Signature: vr,
			ResultCode: SignatureNotCorrect,
			Result: false}
		w.Write(ResponseSuccess(resp).Bytes())
	}

}