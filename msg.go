package main

const(
	VerifySignature int = iota
)

const(
	Sucess int = iota
	MethodNotCorrect
	InterIOErr
	UnmarshalJsonErr
	ActionErr
)

type Request struct {
	Action  int         `json:"action"`
	PayLoad interface{} `json:"payload"`
}

type Response struct {
	Success bool  	      `json:"success"`
	ErrMsg  string        `json:"msg"`
	ErrCode int           `json:"err_code"`
	PayLoad interface{}   `json:"payload"`
}

type VerifyReq struct {
	DID       string `json:"did"` ///public key in string
	Signature string `json:"sig"` ///hex string
	TimeStamp int64 `json:"time_stamp"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type VerifyResp struct {
	Signature *VerifyReq `json:"signature"`
	ResultCode int `json:"result_code"`
	Result bool `json:"result"`
}