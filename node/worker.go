package node

import (
	"context"
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/didchain/didCard-go/account"
	"github.com/didchain/didchain-lightnode-go/protocol"
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

	addr := "0.0.0.0:" + strconv.Itoa(w.port)

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
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.MethodNotCorrect], protocol.MethodNotCorrect).Bytes())
		return
	}

	rbytes,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		w.WriteHeader(200)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.InterIOErr], protocol.InterIOErr).Bytes())
		return
	}

	vr:=&protocol.VerifyReq{}

	req:=&protocol.Request{
		PayLoad: vr,
	}

	if err:=json.Unmarshal(rbytes,req);err!=nil{
		w.WriteHeader(200)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.UnmarshalJsonErr], protocol.UnmarshalJsonErr).Bytes())
		return
	}

	if req.Action != protocol.VerifySignature {
		w.WriteHeader(200)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.ActionErr], protocol.ActionErr).Bytes())
		return
	}

	b:=account.VerifySig(account.ID(vr.DID),base58.Decode(vr.Signature),vr.VerfiyPlainMsg)

	if b{
		w.WriteHeader(200)
		resp:=&protocol.VerifyResp{Signature: vr}
		w.Write(protocol.ResponseSuccess(resp).Bytes())
	}else{
		w.WriteHeader(200)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.SignatureNotCorrect], protocol.SignatureNotCorrect).Bytes())
	}

}