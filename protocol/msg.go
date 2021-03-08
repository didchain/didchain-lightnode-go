package protocol

import "encoding/json"

const(
	VerifySignature int = iota
)

const(
	Sucess int = iota
	SignatureNotCorrect
	InterIOErr
	UnmarshalJsonErr
	ActionErr
	MethodNotCorrect
	UserNotAuthorized
)


var ErrDesc map[int]string =
	map[int]string{
		Sucess:              "success",
		MethodNotCorrect:    "method not correct",
		InterIOErr:          "server read error",
		UnmarshalJsonErr:    "unmarshal json object error",
		ActionErr:           "action not correct",
		SignatureNotCorrect: "signature not correct",
	    }


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


type VerfiyPlainMsg struct {
	DID       string `json:"did"` ///public key in string
	TimeStamp int64 `json:"time_stamp"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type VerifyReq struct {
	Signature string `json:"sig"` ///hex string
	VerfiyPlainMsg
}


type VerifyResp struct {
	Signature *VerifyReq `json:"signature"`
}

func (r *Response)String()  string{
	return string(r.Bytes())
}

func (r *Response)Bytes() []byte  {
	j,_:=json.MarshalIndent(r," ","\t")
	return j
}

func SimpleSuccessResponse() *Response {
	 return &Response{Success: true}
}

func ResponseSuccess(v interface{}) *Response {
	return &Response{Success: true,PayLoad: v}
}

func ResponseError(errMsg string, errCode int) *Response {
	return &Response{ErrCode: errCode,ErrMsg: errMsg}
}

