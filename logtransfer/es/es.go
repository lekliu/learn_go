package es

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/olivere/elastic/v7"
	"net/http"
	"strings"
	"time"
)

type ESChanData struct {
	Topic string `json:"topic"`
	Data  *LogData
}

type LogData struct {
	Data string `json:"data"`
}

// 初始化ES， 准备接收kafka那边发来的数据
var (
	client *elastic.Client
	ch     chan *ESChanData
)

func Init(address string, chanSize int, chanNumber int) (err error) {
	if !strings.HasPrefix(address, "http") {
		address = "https://" + address
	}
	fmt.Println("address: ", address)
	client, err = elastic.NewClient(
		elastic.SetURL(address),
		elastic.SetBasicAuth("elastic", "LU4xOP5+-iNeARhbcK9E"),
		elastic.SetSniff(false),       // 禁用Sniffing，防止客户端自动检测集群节点
		elastic.SetHealthcheck(false), // 禁用健康检查
		elastic.SetScheme("https"),    // 设置为https
		elastic.SetHttpClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 忽略证书验证
				},
			},
		}),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")
	ch = make(chan *ESChanData, chanSize)
	for i := 0; i < chanNumber; i++ {
		go SendToES()
	}
	return err
}

// SendToES 发送数据到ES
func SendToEsChan(topic string, msg *LogData) {
	esChanData := ESChanData{
		Topic: topic,
		Data:  msg,
	}
	ch <- &esChanData
}

func SendToES() error {
	fmt.Println("启动了一个SendToSE goroutine")
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().Index(msg.Topic).BodyJson(msg.Data).Do(context.Background())
			if err != nil {
				fmt.Println("es client send data failed, err: ", err)
			}
			fmt.Printf("ES client send data Topic: %s _Id:%s to index %s, type is %s\n",
				msg.Topic, put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
