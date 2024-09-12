package taillog

import (
	"code.xxliu.com/lekliu/logagent/kafka"
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// TailTask: 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	//为了能实现退出t.run
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path string, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.Init() //根据路径去打开对应的日志
	return
}

// 专门从日志文件收集日志的模块
func (t *TailTask) Init() (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开，日志文件到了一定大小，就会分裂
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个位置开始读
		MustExist: false,                                // 是否必须存在，如果不存在是否报错
		Poll:      true,                                 // Poll for file changes instead of using inotify
	}

	t.instance, err = tail.TailFile(t.path, config)
	fmt.Printf("开始监控文件，path:%v,topic:%v\n", t.path, t.topic)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}

	//当goroutine执行的函数退出的时候，gorotine就结束了
	go t.run() //直接去采休日志发送给Kafka
	return
}

func (t *TailTask) run() {
	tailObj := t.instance
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task %s_%s stop\n", t.path, t.topic)
			return
		// 从tails中一行一行的读取
		case line, ok := <-tailObj.Lines:
			if !ok {
				fmt.Println("tail file close reopen, filename:%s\n", t.path)
				continue
			} else {
				//kafka.SendKafka(t.topic, line.Text)
				//先把日志数据发到一个通道中
				kafka.SendToChan(t.topic, line.Text)
				// kafka那个包中有单独的goroutine 去取日志数据发到kafka
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
