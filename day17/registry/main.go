package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	// 设置 续期 5秒
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	// 将k-v 设置到etcd
	_, err = cli.Put(context.TODO(), "root", "admin", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
	//若想一直有效，设置定期续期
	alive, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c := <-alive
		fmt.Println("c:", c)
	}
}
