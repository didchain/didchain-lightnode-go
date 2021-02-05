package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"time"
)

type Worker struct {
	webserver *http.Server
}

func (w *Worker) process(req Request) *Response {
	panic("stub")
}


func (w *Worker)StartWebDaemon() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ed25519/signatureVerify", blsVerify)

	addr := "127.0.0.1:50999"

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

func blsVerify(w http.ResponseWriter,r *http.Request)  {
	if r.Method != "POST"{
		w.WriteHeader(500)
		w.Write([]byte("method is not correct"))
		return
	}

	rbytes,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		w.WriteHeader(200)
		w.Write([]byte("sign server internal error"))
		return
	}

	req:=&Request{
		PayLoad: &VerifyReq{},
	}

	if err:=json.Unmarshal(rbytes,req);err!=nil{
		w.WriteHeader(200)
		w.Write([]byte(err.Error()))
		return
	}

	if req.Action != VerifySignature{
		w.WriteHeader(200)
		w.Write([]byte(""))
		return
	}

}