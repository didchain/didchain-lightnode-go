package loginUam

import (
	"github.com/didchain/didchain-lightnode-go/protocol"

)

type SessionInfo struct {
	Signature *protocol.UAMSignature
	UserDesc *protocol.UserDesc
}
