package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

//- 定义MemorySessionMgr对象（字段：存放所有session的map，读写锁)
//- 构造函数
//- Init(),
//- CreateSession()，创建一个新的Session
//- GetSession()

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	sm := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 100),
		rwLock:     sync.RWMutex{},
	}
	return sm
}

// Init
func (msm *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

// 创建 Session
func (msm *MemorySessionMgr) CreateSession() (session Session, err error) {
	msm.rwLock.Lock()
	defer msm.rwLock.Unlock()
	//用UUID生成sesionid
	sessionId := uuid.NewV4().String()
	//创建单个session
	session = NewMemorySession(sessionId)
	msm.rwLock.Lock()
	defer msm.rwLock.Unlock()
	msm.sessionMap[sessionId] = session
	return
}

// getSession
func (msm *MemorySessionMgr) GetSession(sessionId string) (s Session, err error) {
	msm.rwLock.RLock()
	defer msm.rwLock.RUnlock()
	var ok bool
	s, ok = msm.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not found")
		return
	}
	return
}

//go get github.com/satori/go.uuid
