package session

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/kprc/nbsnetwork/tools"
	"math/rand"
	"sync"
	"time"
)

type SessionDesc struct {
	IsVerify       bool
	LastAccessTime int64
}

const (
	RandBytesCount int = 16
)

var (
	sessionTimeout int = 600 //10 minutes
	quit           chan struct{}
	wg             sync.WaitGroup
	accessSession  map[[RandBytesCount]byte]*SessionDesc
)

func init() {
	quit = make(chan struct{})
	accessSession = make(map[[RandBytesCount]byte]*SessionDesc)
}

func NewSession() ([RandBytesCount]byte, *SessionDesc) {
	var randbytes [RandBytesCount]byte

	for {
		rand.Seed(tools.GetNowMsTime())
		n, err := rand.Read(randbytes[:])
		if err != nil || RandBytesCount != n {
			continue
		}
		break
	}

	sd := &SessionDesc{LastAccessTime: tools.GetNowMsTime()}

	accessSession[randbytes] = sd

	return randbytes, sd

}

func NewSession2() (string, *SessionDesc) {
	b, s := NewSession()

	return base58.Encode(b[:]), s
}

func IsSession(k [RandBytesCount]byte) bool {
	if _, ok := accessSession[k]; !ok {
		return false
	}

	return true
}

func IsSessionBase58(k string) bool {
	kb := base58.Decode(k)

	if len(kb) != RandBytesCount {
		return false
	}

	var key [RandBytesCount]byte
	copy(key[:], kb)

	return IsSession(key)
}

func IsValid(k [RandBytesCount]byte) bool {
	if s, ok := accessSession[k]; !ok {
		return false
	} else {
		if !s.IsVerify {
			return false
		}

		if tools.GetNowMsTime()-s.LastAccessTime > (int64(sessionTimeout) * 1000) {
			return false
		}

		s.LastAccessTime = tools.GetNowMsTime()

		return true
	}
}

func IsValidBase58(k string) bool {
	kb := base58.Decode(k)

	if len(kb) != RandBytesCount {
		return false
	}

	var key [RandBytesCount]byte
	copy(key[:], kb)

	return IsValid(key)
}

func SessionActiveBase58(k string) {
	kb := base58.Decode(k)

	if len(kb) != RandBytesCount {
		return
	}

	var key [RandBytesCount]byte
	copy(key[:], kb)

	if v, ok := accessSession[key]; !ok {
		return
	} else {
		v.IsVerify = true
		v.LastAccessTime = tools.GetNowMsTime()
	}

}

func StartTimeOut() {
	wg.Add(1)
	for {

		select {
		case <-quit:
			wg.Done()
			return
		default:

		}

		now := tools.GetNowMsTime()

		var ks [][RandBytesCount]byte
		for k, v := range accessSession {
			if now-v.LastAccessTime > (int64(sessionTimeout) * 1000) {
				ks = append(ks, k)
			}
		}

		for i := 0; i < len(ks); i++ {
			delete(accessSession, ks[i])
		}

		time.Sleep(time.Second)
	}
}

func StopTimeOut() {
	quit <- struct{}{}

	wg.Wait()
}
