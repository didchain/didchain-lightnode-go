package node

import (
	"context"
	"github.com/didchain/didchain-lightnode-go/user/storage"
	"github.com/didchain/didchain-lightnode-go/user/webapi"
	"log"
	"net/http"
	"strconv"

	"time"
)

type Worker struct {
	port int
	webserver *http.Server
	storage *storage.Storage
}


func (w *Worker)StartWebDaemon() {
	mux := http.NewServeMux()

	userapi:=webapi.NewUserAPI(w.storage)

	mux.HandleFunc("/ed25519/signatureVerify", userapi.Ed25519Verify)
	mux.HandleFunc("/api/user/add",userapi.AddUser)
	mux.HandleFunc("/api/user/del",userapi.DelUser)
	mux.HandleFunc("/api/user/count",userapi.UserCount)
	mux.HandleFunc("/api/user/listUser",userapi.ListUser)
	mux.HandleFunc("/api/user/listUser4Add",userapi.ListUnAuthorizeUser)
	mux.HandleFunc("/api/auth/token",webapi.AccessToken)
	mux.HandleFunc("api/auth/verify",webapi.SigVerify)

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
