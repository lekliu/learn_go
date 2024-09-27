package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

//- 定义RedisSessionMgr对象（字段：存放所有session的map，读写锁)
//- 构造函数
//- Init(),
//- CreateSession()，创建一个新的Session
//- GetSession()

type RedisSessionMgr struct {
	//redis地址
	addr string
	// 密码
	passwd string
	//连接池
	pool       *redis.Pool
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

// 构造函数
func NewRedisSessionMgr() *RedisSessionMgr {
	sm := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 100),
		rwLock:     sync.RWMutex{},
	}
	return sm
}

// Init
func (rm *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	//若有其它参数
	if len(options) > 0 {
		rm.passwd = options[0]
	}
	// 创建连接池
	myPool(addr, rm.passwd)
	rm.addr = addr
	return
}

func myPool(addr, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if len(passwd) > 0 {
				if _, err := conn.Do("AUTH", passwd); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		//连接测试，开发时写，上线会注释掉
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

// 创建 Session
func (rm *RedisSessionMgr) CreateSession() (session Session, err error) {
	rm.rwLock.Lock()
	defer rm.rwLock.Unlock()
	//用UUID生成sesionid
	sessionId := uuid.NewV4().String()
	//创建单个session
	session = NewRedisSession(sessionId, rm.pool)
	rm.rwLock.Lock()
	defer rm.rwLock.Unlock()
	rm.sessionMap[sessionId] = session
	return
}

// getSession
func (rm *RedisSessionMgr) GetSession(sessionId string) (s Session, err error) {
	rm.rwLock.RLock()
	defer rm.rwLock.RUnlock()
	var ok bool
	s, ok = rm.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not found")
		return
	}
	return
}

//go get github.com/satori/go.uuid
