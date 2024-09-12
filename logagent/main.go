package main

import (
	"code.xxliu.com/lekliu/logagent/conf"
	"code.xxliu.com/lekliu/logagent/etcd"
	"code.xxliu.com/lekliu/logagent/kafka"
	"code.xxliu.com/lekliu/logagent/taillog"
	"code.xxliu.com/lekliu/logagent/utils"
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

var (
	cfg = new(conf.AppConfig)
)

func main() {
	//0.加载配置文件
	//cfg, err := ini.Load("./conf/config.ini")
	//address := cfg.Section("kafka").Key("address").String()
	//topic := cfg.Section("kafka").Key("topic").String()
	//path := cfg.Section("taillog").Key("path").String()
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini faild,err:%v\n", err)
		return
	}
	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka faild,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")

	//2. 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd faild,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")
	// 为了实现每个LogAgent都拉取自己独有的配置，所以要以自己的IP地址作为区分
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	// 2.1 从etcd中获取日志收集项的配置信息
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Collect_log_key, ipStr)
	fmt.Println("the config key of etcd is ", etcdConfKey)
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("get etcd config faild,err:%v\n", err)
	}
	fmt.Printf("get etcd config success，%v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%d,value:%v\n", index, value)
	}

	// 3. 收集日志发往Kafka
	taillog.Init(logEntryConf)

	// 2.2 派一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现加载配置）
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, taillog.NewConfChan())
	wg.Wait()
}
