package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type LogData struct {
	topic string
	data  string
}

// 专门往kafka写日志的模块
var (
	client      sarama.SyncProducer //声明一个全局的连接kafka的生产者client
	logDataChan chan *LogData
)

// Init 初化化client
func Init(addrs []string, chanMaxSize int) (err error) {
	config := sarama.NewConfig()

	// tailf包使用，发送完数据需要 leader 和 follow都确定
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的消息将在 success channel返回
	config.Producer.Return.Successes = true

	// 连接kafka，可以连接一个集群
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err: ", err)
	}
	logDataChan = make(chan *LogData, chanMaxSize)
	//开启后台的goroutine从通道中取数据发往kafka
	go sendKafka()
	return
}

// 给外部暴露的一个函数，该函数只把日志数据发送到一个内部的channel中
func SendToChan(topic string, data string) {
	logData := &LogData{
		topic: topic,
		data:  data}
	logDataChan <- logData
}

// 真正往Kafka发送日志的函数
func sendKafka() (err error) {
	var pid int32
	var offset int64
	for {
		select {
		case logData := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = logData.topic
			msg.Value = sarama.StringEncoder(logData.data)
			//发送到kafka
			pid, offset, err = client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v \n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
	return
}
