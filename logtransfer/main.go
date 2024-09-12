package main

import (
	"code.xxliu.com/lekliu/logtransfer/conf"
	"code.xxliu.com/lekliu/logtransfer/es"
	"code.xxliu.com/lekliu/logtransfer/kafka"
	"fmt"
	"gopkg.in/ini.v1"
)

// log transfer

var (
	cfg = new(conf.LogTransferConfig)
)

// 将日志数据从kafka取出来，发往ES
func main() {
	//0. 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load config faild,err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", cfg)

	// 1.从Kafka取日志数据
	// 1.1 初始化一个ES连接的client
	// 1。2 对外提供一个往ES写入数据的函数
	err = es.Init(cfg.EsConf.Address, cfg.EsConf.ChanSize, cfg.EsConf.ChanNumber)
	if err != nil {
		fmt.Printf(" init es faild, err:%v\n", err)
		return
	}
	fmt.Println("init ES client success.")

	// 2. 初始化kafka
	// 2.1 连接kafka，创建分区的消费者
	// 2.2 每个分区的消费者分别取出数据，通过sendToES()将数据发往ES
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Printf("kafka consumer init faild, err:%v\n", err)
		return
	}
	select {}
}
