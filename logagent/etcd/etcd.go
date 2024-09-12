package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// 需要收集的日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`  //日志存放的路径
	Topic string `json:"topic"` //日志要发往Kafka的哪个Topic
}

var (
	cli *clientv3.Client
)

// 初始化 etcd的方法
func Init(addrs string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addrs},
		DialTimeout: timeout,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	return
}

// 从ETCD中根据Key获取配置项
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf(" unmarshal etcd value failed, err:%v\n", err)
			return
		}
	}
	return
}

// etcd watch
func WatchConf(key string, newConfChan chan<- []*LogEntry) {
	rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			//通知 tailog.tskMgr
			// 1. 先判断操作的类开
			var newConf []*LogEntry
			if ev.Type != clientv3.EventTypeDelete {
				//如果是删除操作，手动传递一个空的配置项，否则json解析错误
				err := json.Unmarshal(ev.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal etcd value failed, err:%v\n", err)
					continue
				}
			}
			fmt.Printf(" get new conf:%v\n", newConf)
			newConfChan <- newConf
		}
	}
}
