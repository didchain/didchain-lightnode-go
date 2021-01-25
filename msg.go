package main

type Request struct {
	Action  int         `json:"action"`
	PayLoad interface{} `json:"payload"`
}

type Response struct {
	Success bool        `json:"success"`
	Msg     bool        `json:"msg"`
	PayLoad interface{} `json:"payload"`
}

type VerifyReq struct {
	DID       string `json:"did"` ///public key in string
	Signature string `json:"sig"` ///hex string
	Msg       string `json:"msg"` ///hex string
}
