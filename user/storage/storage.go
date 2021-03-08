package storage

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/btcsuite/goleveldb/leveldb/util"
	"sync"
)

const(
	UserHead string = "didchain_user_%s"
	UserListBegin string = "didchain_user_did111111111111111111111111111111111"
	UserListEnd   string = "didchain_user_didzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
)

type Storage struct {
	lock sync.RWMutex
	db *leveldb.DB
}

func NewStorage(db *leveldb.DB) *Storage  {
	return &Storage{db: db}
}


func didUserKey(did string) string  {
	return fmt.Sprintf(UserHead,did)
}

func (s *Storage)AddUser(did string, v interface{}) error {
	wo := &opt.WriteOptions{
		Sync: true,
	}

	j,_:=json.Marshal(v)

	s.lock.Lock()
	defer s.lock.Unlock()

	if err:=s.db.Put([]byte(didUserKey(did)),j,wo);err!=nil{
		return err
	}

	return nil
}

func (s *Storage)DelUser(did string)   {
	wo := &opt.WriteOptions{
		Sync: true,
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if err:=s.db.Delete([]byte(didUserKey(did)),wo);err!=nil{
		fmt.Println("Storage delete",did,"failed")
	}
}

func (s *Storage)FindUser(did string) string {

	s.lock.Lock()
	defer s.lock.Unlock()

	v,err:=s.db.Get([]byte(didUserKey(did)),nil)
	if err!=nil{
		return ""
	}

	return string(v)
}

func (s *Storage)ListAll() []string  {
	s.lock.RLock()
	defer s.lock.RUnlock()

	r:=&util.Range{Start: []byte(UserListBegin),Limit: []byte(UserListEnd)}
	iter:=s.db.NewIterator(r,nil)
	defer iter.Release()

	var vs []string

	for iter.Next(){
		vs = append(vs,string(iter.Value()))
	}

	return vs
}


func (s *Storage)ListAllValue(convert func(data []byte) interface{}) []interface{}  {
	s.lock.RLock()
	defer s.lock.RUnlock()

	r:=&util.Range{Start: []byte(UserListBegin),Limit: []byte(UserListEnd)}
	iter:=s.db.NewIterator(r,nil)
	defer iter.Release()

	var vs []interface{}

	for iter.Next(){
		vs = append(vs,convert(iter.Value()))
	}

	return vs
}

func (s *Storage)ListAllValue2(convert func(data []byte) interface{}, begin, count int) []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	r:=&util.Range{Start: []byte(UserListBegin),Limit: []byte(UserListEnd)}
	iter:=s.db.NewIterator(r,nil)
	defer iter.Release()

	var (
		vs []interface{}
		cnt int
	)

	for iter.Next(){
		if begin<=cnt{
			vs = append(vs,convert(iter.Value()))
		}

		if cnt - begin +1 > count{
			break
		}

		cnt ++

	}

	return vs
}

