package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/olivere/elastic/v7"
	"net/http"
)

type student struct {
	Name    string `json:"name""`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func (s *student) run() *student {
	fmt.Printf("%s在跑...", s.Name)
	return s
}

func (s *student) wang() *student {
	fmt.Printf("%s在汪汪汪的叫...", s.Name)
	return s
}

func main() {
	//luminghui := student{
	//	Name:    "luminghui",
	//	Age:     2000,
	//	Married: true,
	//}
	//luminghui.run()
	//luminghui.wang()
	//luminghui.run().wang()

	// 1. 初始化连接， 得到一个Client
	client, err := elastic.NewClient(
		elastic.SetURL("https://127.0.0.1:9200"),
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
		//handle error
		panic(err)
	}
	fmt.Println("connect to es success")
	p1 := student{Name: "Tom", Age: 22, Married: false}
	// 链式操作
	put1, err := client.Index().Index("student").BodyJson(p1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed student %s to index %s, type is %s\n", put1.Id, put1.Index, put1.Type)

}
