package protocol

type AuthContent struct {
	AuthUrl     string `json:"auth_url"`
	RandomToken string `json:"random_token"`
}

type UAMSignatureContent struct {
	AuthUrl     string `json:"auth_url"`
	RandomToken string `json:"random_token"`
	DID         string `json:"did"`
}

type UAMExtData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UAMSignature struct {
	Content   *UAMSignatureContent `json:"content"`
	ExtData   *UAMExtData          `json:"ext_data"`
	Signature string               `json:"sig"`
}

type UserDesc struct {
	Name         string `json:"name"`
	UnitName     string `json:"unit_name"`
	SerialNumber string `json:"serial_number"`
	Did          string `json:"did"`
}

type UAMCheck struct {
	RedirUrl  string        `json:"redir_url"`
	Signature *UAMSignature `json:"signature"`
	UserDesc  *UserDesc     `json:"user_desc"`
}

const (
	Success      int = 0
	SigNatureErr int = 1
	UserNotFound int = 2
	InternalErr  int = 999
)

const (
	SuccessMsg      string = "success"
	SigNatureErrMsg string = "signature not correct"
	UserNotFoundMsg string = "user not found"
	InternalErrMsg  string = "server internal error"
)

type UAMResponse struct {
	ResultCode int         `json:"result_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
