package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/didchain/didchain-lightnode-go/user/session"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func AccessToken(w http.ResponseWriter, r *http.Request) {
	randbytes, _ := session.NewSession()

	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", base58.Encode(randbytes[:]))
}

var logger, _ = logging.GetLogger("webserver")

type AccessSig struct {
	Sig        string `json:"sig"`
	AccesToken string `json:"acces_token"`
}

type ValidSigResult struct {
	ResultCode  int    `json:"result_code"` //0 success, 1 session not found, 2 signature not correct, 3 other error
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

func (ua *UserAPI) SigVerify(w http.ResponseWriter, r *http.Request) {

	vsr := ua.doSigVerify(r)

	j, _ := json.Marshal(*vsr)

	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", string(j))

}

func (ua *UserAPI) doSigVerify(r *http.Request) *ValidSigResult {
	vsr := &ValidSigResult{}

	if r.Method != "POST" {
		vsr.ResultCode = 3
		vsr.Message = "must a post request"
		return vsr
	}
	fmt.Println(1)
	var (
		content []byte
		err     error
	)
	if content, err = ioutil.ReadAll(r.Body); err != nil {
		vsr.ResultCode = 3
		vsr.Message = "read http body error"
		return vsr
	}
	fmt.Println(2)

	as := &AccessSig{}
	err = json.Unmarshal(content, as)
	if err != nil {
		vsr.ResultCode = 3
		vsr.Message = "json string error"
		return vsr
	}
	fmt.Println(3)
	if !session.IsSessionBase58(as.AccesToken) {
		vsr.ResultCode = 1
		vsr.Message = "token not found"
		return vsr
	}
	fmt.Println(4)
	bsig := base58.Decode(as.Sig)
	if len(bsig) == 0 {
		vsr.ResultCode = 2
		vsr.Message = "signature must encode by base58"
	} else {
		to := "\x19Ethereum Signed Message:\n"
		to += strconv.Itoa(len(as.AccesToken))
		to += as.AccesToken

		logger.Info("message:", to)
		logger.Info("sig:", as.Sig)

		//if !microchain.ChainInst().VerifySig([]byte(to), base58.Decode(as.Sig)) {
		//	vsr.ResultCode = 2
		//	vsr.Message = "signature not correct"
		//}
		fmt.Println(5)
		if !ua.verify(to, as.Sig) {
			vsr.ResultCode = 2
			vsr.Message = "signature not correct"
		}
		fmt.Println(6)
	}

	if vsr.ResultCode == 0 {
		vsr.Message = "success"
		vsr.AccessToken = as.AccesToken
		fmt.Println(7)
		session.SessionActiveBase58(as.AccesToken)
		fmt.Println(8)
	}

	return vsr
}

func (ua *UserAPI) verify(message string, sigstr string) bool {

	return true

	hash := crypto.Keccak256([]byte(message))
	sig := base58.Decode(sigstr)
	idx := len(sig) - 1
	fmt.Println(11)
	if sig[idx] > 1 {
		sig[idx] = byte(sig[idx]) - 0x1b
	}

	recoveredPub, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		//fmt.Println("1", err)
		return false
	}
	fmt.Println(12)

	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	raddr := recoveredAddr.String()

	fmt.Println(13)

	addrs := ua.admin.ListUser()

	fmt.Println(raddr)

	for _, addr := range addrs {
		if strings.ToLower(raddr) == strings.ToLower(addr) {
			return true
		}
	}

	fmt.Println(15)

	return false

}
