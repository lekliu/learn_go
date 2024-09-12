package kafka

import (
	"code.xxliu.com/lekliu/logtransfer/es"
	"fmt"
	"github.com/Shopify/sarama"
)

// Init 初始化
func Init(addrs []string, topic string) (err error) {
	var consumer sarama.Consumer
	var partitionList []int32
	var pc sarama.PartitionConsumer

	consumer, err = sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err = consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println("分区列表：", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				//发往ES
				ld := es.LogData{
					Data: string(msg.Value),
				}

				es.SendToEsChan(topic, &ld) //函数调函数
				//优化一下；直接放到一个chan中
			}
		}(pc)
	}
	return
}
