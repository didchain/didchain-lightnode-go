package loginUam

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/didchain/didchain-lightnode-go/config"
	"github.com/kprc/nbsnetwork/tools"
	"math/rand"
	"sync"
	"time"
)

type SessID [32]byte

type SessionIntf interface {
	NewSession() SessID
	AddSession(sessid SessID, v interface{}) //add or update session
	FindSession(sessid SessID) interface{}
	DelSession(sessid SessID)
	TimeOut()
}

type Session struct {
	lastAccessTime int64
	id             SessID
	v              interface{}
}

type SessStorage struct {
	lock       sync.RWMutex
	cfg        *config.NodeConfig
	sessionSet map[SessID]*Session
	quit       chan struct{}
}

func NewSessStorage(cfg *config.NodeConfig) *SessStorage {
	ss := &SessStorage{
		cfg:        cfg,
		sessionSet: make(map[SessID]*Session),
		quit:       make(chan struct{}, 1),
	}

	return ss
}

func Str2SessID(sess string) SessID {
	var id SessID

	idbytes := base58.Decode(sess)

	copy(id[:], idbytes)
	return id
}

func (ss *SessStorage) NewSession() SessID {
	var id SessID

	for {
		rand.Seed(tools.GetNowMsTime())
		n, err := rand.Read(id[:])
		if err != nil || len(id) != n {
			continue
		}
		break
	}

	return id
}

func (ss *SessStorage) AddSession(sessid SessID, v interface{}) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	sess := &Session{
		lastAccessTime: tools.GetNowMsTime(),
		id:             sessid,
		v:              v,
	}

	ss.sessionSet[sessid] = sess

	return
}

func (ss *SessStorage) DelSession(sessid SessID) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	delete(ss.sessionSet, sessid)

}

func (ss *SessStorage) FindSession(sessid SessID) interface{} {
	ss.lock.RLock()
	defer ss.lock.RUnlock()

	sess, ok := ss.sessionSet[sessid]
	if !ok {
		return nil
	}
	sess.lastAccessTime = tools.GetNowMsTime()
	return sess.v
}

func (ss *SessStorage) TimeOut() {
	var lastTime int64

	for {
		select {
		case <-ss.quit:
			return
		default:
			if tools.GetNowMsTime()-lastTime > 16000 {
				ss.timeOut()
			}
		}
		time.Sleep(time.Second * 1)
	}
}

func (ss *SessStorage) timeOut() {
	dels := ss.findTimeout()

	for i := 0; i < len(dels); i++ {
		ss.DelSession(dels[i])
	}
}

func (ss *SessStorage) findTimeout() []SessID {
	var dels []SessID

	ss.lock.RLock()
	defer ss.lock.RUnlock()

	now := tools.GetNowMsTime()

	for k, v := range ss.sessionSet {
		if now-v.lastAccessTime >= ss.cfg.SessiontimeOut {
			dels = append(dels, k)
		}
	}

	return dels
}
