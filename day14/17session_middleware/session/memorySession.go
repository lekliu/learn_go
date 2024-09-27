package session

import (
	"errors"
	"sync"
)

// * 定义MemorySession对象（字段：sessionId、存kv的map，读写锁）
// * 构造函数，为了获取对象
// * Set()
// * Get()
// * Del()
// * Save()
// 对象
type MemorySession struct {
	sessionId string
	//存kv
	data   map[string]interface{}
	rwLock sync.RWMutex
}

// 构造函数
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 10),
		rwLock:    sync.RWMutex{},
	}
	return s
}

// set
func (s *MemorySession) Set(key string, value interface{}) (err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.data[key] = value
	return
}

// Get
func (s *MemorySession) Get(key string) (value interface{}, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	value, ok := s.data[key]
	if !ok {
		err = errors.New("key(" + key + ") is not in session")
		return
	}
	return
}

// Del
func (s *MemorySession) Delete(key string) (err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	delete(s.data, key)
	return
}

// Save
func (s *MemorySession) Save() (err error) {
	return
}
