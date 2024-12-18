package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// etcd client put/get demo
// use etcd/clientv3

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"c:/xxx/nginx.log","topic":"web_log"},{"path":"c:/xxx/redis.log","topic":"redis_log"}]`
	//value := `[{"path":"c:/xxx/nginx.log","topic":"web_log"},{"path":"c:/xxx/redis.log","topic":"redis_log"},{"path":"c:/xxx/mysql.log","topic":"mysql_log"}]`
	//value := `[{"path":"c:/xxx/nginx.log","topic":"web_log"}]`
	ip, err := GetOutboundIP()
	_, err = cli.Put(ctx, "/logagent/"+ip+"/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
