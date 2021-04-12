package storage

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/btcsuite/goleveldb/leveldb/util"
	"github.com/didchain/didCard-go/account"
	"strings"
	"sync"
)

const (
	UserHead      string = "didchain_user_%s"
	UserListBegin string = "didchain_user_did"
	UserListEnd   string = "didchain_user_didzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
)

const(
	hostHead string = "didchain_host_%s"
	hostPatternBegin string = "didchain_host_"
	hostPatternEnd string = "didchain_host_~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
)


type Storage struct {
	lock sync.RWMutex
	db   *leveldb.DB
}

func NewStorage(db *leveldb.DB) *Storage {
	return &Storage{db: db}
}

func didUserKey(did string) string {
	return fmt.Sprintf(UserHead, did)
}

func (s *Storage) AddUser(did string, v interface{}) error {
	wo := &opt.WriteOptions{
		Sync: true,
	}

	j, _ := json.Marshal(v)

	fmt.Println("add user", string(j))

	if _, err := account.ConvertToID(did); err != nil {
		return err
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if err := s.db.Put([]byte(didUserKey(did)), j, wo); err != nil {
		fmt.Println("add user", err)
		return err
	}

	return nil
}

func (s *Storage) DelUser(did string) {
	wo := &opt.WriteOptions{
		Sync: true,
	}

	fmt.Println("delete user", did)

	s.lock.Lock()
	defer s.lock.Unlock()

	if err := s.db.Delete([]byte(didUserKey(did)), wo); err != nil {
		fmt.Println("Storage delete", did, "failed")
	}
}

func (s *Storage) FindUser(did string) string {

	s.lock.Lock()
	defer s.lock.Unlock()

	v, err := s.db.Get([]byte(didUserKey(did)), nil)
	if err != nil {
		return ""
	}

	fmt.Println("find user", string(v))

	return string(v)
}

func (s *Storage) ListAll() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	r := &util.Range{Start: []byte(UserListBegin), Limit: []byte(UserListEnd)}
	iter := s.db.NewIterator(r, nil)
	defer iter.Release()

	var vs []string

	for iter.Next() {
		fmt.Println("listall", string(iter.Key()))
		vs = append(vs, string(iter.Value()))
	}

	return vs
}

func (s *Storage) Count() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	r := &util.Range{Start: []byte(UserListBegin), Limit: []byte(UserListEnd)}
	iter := s.db.NewIterator(r, nil)
	defer iter.Release()

	var count int

	for iter.Next() {
		count++
	}

	return count
}

func (s *Storage) ListAllValue(convert func(data []byte) interface{}) []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	r := &util.Range{Start: []byte(UserListBegin), Limit: []byte(UserListEnd)}
	iter := s.db.NewIterator(r, nil)
	defer iter.Release()

	var vs []interface{}

	for iter.Next() {
		vs = append(vs, convert(iter.Value()))
	}

	return vs
}

func (s *Storage) ListAllValue2(convert func(data []byte) interface{}, begin, count int) []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	r := &util.Range{Start: []byte(UserListBegin), Limit: []byte(UserListEnd)}
	iter := s.db.NewIterator(r, nil)
	defer iter.Release()

	var (
		vs  []interface{}
		cnt int
	)

	end := begin + count

	for iter.Next() {

		fmt.Println("listall value2", iter.Key())

		if cnt >= begin && cnt < end {
			vs = append(vs, convert(iter.Value()))
		}

		cnt++

		if cnt >= end {
			break
		}

	}

	return vs
}

func (s *Storage)AddHost(host string) error  {
	bas64host := base64.StdEncoding.EncodeToString([]byte(host))

	s.lock.Lock()
	defer s.lock.Unlock()

	key:=fmt.Sprintf(hostHead,bas64host)
	wo:=&opt.WriteOptions{
		Sync: true,
	}
	//log.Println(key)
	return s.db.Put([]byte(key),[]byte("v"),wo)
}

func (s *Storage)DelHost(host string)   {
	bas64host := base64.StdEncoding.EncodeToString([]byte(host))

	s.lock.Lock()
	defer s.lock.Unlock()

	key:=fmt.Sprintf(hostHead,bas64host)
	wo:=&opt.WriteOptions{
		Sync: true,
	}

	s.db.Delete([]byte(key),wo)
}

func (s *Storage)ListAllHost() ([]string)  {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var host []string

	r := &util.Range{Start: []byte(hostPatternBegin), Limit: []byte(hostPatternEnd)}
	iter := s.db.NewIterator(r, nil)
	defer iter.Release()

	for iter.Next() {
		//log.Println("key is",string(iter.Key()))
		ks:=strings.Split(string(iter.Key()),"_")
		if len(ks) != 3{
			continue
		}

		base64host:=ks[2]
		//log.Println("base64host is",base64host)
		h,_:=base64.StdEncoding.DecodeString(base64host)
		//log.Println("host is",h)
		host = append(host,string(h))
	}

	return host

}