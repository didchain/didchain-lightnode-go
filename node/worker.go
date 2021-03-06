package node

import (
	"context"
	"github.com/didchain/didchain-lightnode-go/config"
	"github.com/didchain/didchain-lightnode-go/user/session"
	"github.com/didchain/didchain-lightnode-go/user/storage"
	"github.com/didchain/didchain-lightnode-go/user/webapi"
	"github.com/didchain/didchain-lightnode-go/webpages/webfs"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"log"
	"net/http"
	"strconv"

	"time"
)

type Worker struct {
	port      int
	webserver *http.Server
	storage   *storage.Storage
	admin     *config.AdminUser
}

func (w *Worker) StartWebDaemon() {
	mux := http.NewServeMux()

	userapi := webapi.NewUserAPI(w.storage, w.admin)

	mux.HandleFunc("/ed25519/signatureVerify", userapi.Ed25519Verify)
	mux.HandleFunc("/api/user/add", userapi.AddUser)
	mux.HandleFunc("/api/user/del", userapi.DelUser)
	mux.HandleFunc("/api/user/count", userapi.UserCount)
	mux.HandleFunc("/api/user/listUser", userapi.ListUser)
	mux.HandleFunc("/api/user/listUser4Add", userapi.ListUnAuthorizeUser)
	mux.HandleFunc("/api/auth/token", webapi.AccessToken)
	mux.HandleFunc("/api/auth/verify", userapi.SigVerify)

	wfs := assetfs.AssetFS{Asset: webfs.Asset, AssetDir: webfs.AssetDir, AssetInfo: webfs.AssetInfo, Prefix: "webpages/html/dist"}

	mux.Handle("/", http.FileServer(&wfs))

	addr := "0.0.0.0:" + strconv.Itoa(w.port)

	log.Println("Web Server Start at", addr)

	w.webserver = &http.Server{Addr: addr, Handler: mux}

	go session.StartTimeOut()

	go log.Fatal(w.webserver.ListenAndServe())
}

func (w *Worker) StopWebDaemon() {
	if w.webserver == nil {
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	w.webserver.Shutdown(ctx)

	session.StopTimeOut()

	w.webserver = nil

	log.Println("Web Server Stopped")
}
