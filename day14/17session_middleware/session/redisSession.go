package session

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"sync"
)

// * 定义redisSession对象（字段：sessionId、存kv的map，读写锁）
// * 构造函数，为了获取对象
// * Set()
// * Get()
// * Del()
// * Save()

// 用常量去定义状态
const (
	//内存数据没有变化
	SessionFlagNone = iota
	// 有变化
	SessionFlagModify
)

// 对象
type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	// 设置session，可以先放在内存的map中
	// 批量导入redis, 提升性能
	sessionMap map[string]interface{}
	//读写锁
	rwLock sync.RWMutex
	// 记录内存中的map是否被操作
	flag int
}

// 构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}, 16),
		rwLock:     sync.RWMutex{},
		flag:       SessionFlagNone,
	}
	return s
}

// set Session存储到内存中的map
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.sessionMap[key] = value
	//标记记录
	r.flag = SessionFlagModify
	return
}

// Save session在搞到 redis
func (r *RedisSession) Save() (err error) {
	r.rwLock.Lock()
	//defer r.rwLock.Unlock()
	//如果数扰没变，不需要存
	if r.flag != SessionFlagModify {
		return
	}
	//内存中的sessionMap进行序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return err
	}
	//获取redis连接
	conn := r.pool.Get()
	//defer conn.Close()
	_, err = conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return err
	}
	r.flag = SessionFlagNone
	return
}

// Get
func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	// 先判断内存
	result = r.sessionMap[key]
	return
}

// 从redis里再次加载
func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	// 转字符串
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	//取到的东西，反序列化到内存map
	json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return
}

// Del
func (r *RedisSession) Delete(key string) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)

	return
}

//go get "github.com/garyburd/redigo/redis"
