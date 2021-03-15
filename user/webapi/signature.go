package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/didchain/didCard-go/account"
	"github.com/didchain/didchain-lightnode-go/protocol"
	"github.com/kprc/nbsnetwork/tools"
	"io/ioutil"
	"net/http"
)

func (ua *UserAPI) Ed25519Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(500)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.MethodNotCorrect], protocol.MethodNotCorrect).Bytes())
		return
	}

	rbytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(200)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.InterIOErr], protocol.InterIOErr).Bytes())
		return
	}

	vr := &protocol.VerifyReq{}

	req := &protocol.Request{
		PayLoad: vr,
	}

	if err := json.Unmarshal(rbytes, req); err != nil {
		w.WriteHeader(200)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.UnmarshalJsonErr], protocol.UnmarshalJsonErr).Bytes())
		fmt.Println(err, string(rbytes))
		return
	}

	if req.Action != protocol.VerifySignature {
		w.WriteHeader(200)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.ActionErr], protocol.ActionErr).Bytes())
		return
	}

	b := account.VerifySig(account.ID(vr.DID), base58.Decode(vr.Signature), vr.VerfiyPlainMsg)

	if b {

		if v := ua.sdb.FindUser(vr.DID); v == "" {
			glist4add.add(vr.DID, tools.GetNowMsTime())

			w.WriteHeader(200)
			w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.UserNotAuthorized], protocol.UserNotAuthorized).Bytes())
			return
		}

		w.WriteHeader(200)
		resp := &protocol.VerifyResp{Signature: vr}
		w.Write(protocol.ResponseSuccess(resp).Bytes())
	} else {

		w.WriteHeader(200)
		w.Write(protocol.ResponseError(protocol.ErrDesc[protocol.SignatureNotCorrect], protocol.SignatureNotCorrect).Bytes())
	}

}
