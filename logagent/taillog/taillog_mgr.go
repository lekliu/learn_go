package taillog

import (
	"code.xxliu.com/lekliu/logagent/etcd"
	"fmt"
	"time"
)

var tskMgr *TailLogMgr

// tailTask 管理者
type TailLogMgr struct {
	logEntryList []*etcd.LogEntry
	tskMap       map[string]*TailTask
	newConfChan  chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &TailLogMgr{
		logEntryList: logEntryConf, //把当前的日志收集项配置信息保存起来
		tskMap:       make(map[string]*TailTask, 16),
		newConfChan:  make(chan []*etcd.LogEntry), //无缓冲区的通道
	}
	for _, logEntryConf := range logEntryConf {
		fmt.Printf("conf:%v\n", logEntryConf.Path)
		// 初始化的时候，起了多少个tailTask都要记下来，为了后续判断方便
		tailTask := NewTailTask(logEntryConf.Path, logEntryConf.Topic)
		mk := fmt.Sprintf("%s_%s", logEntryConf.Path, logEntryConf.Topic)
		tskMgr.tskMap[mk] = tailTask
	}
	go tskMgr.run()
}

// 监听自己的newConfChan,有了新的配置过来之后就做对应的处理

func (t *TailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, logEntry := range newConf {
				mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 1.配置新增
					tailTask := NewTailTask(logEntry.Path, logEntry.Topic)
					mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
					tskMgr.tskMap[mk] = tailTask
				}
			}
			// 2.配置删除
			// 找出原来 t.logEntry有，但是新的newConf中没有的，删除
			for _, logEntry := range t.logEntryList {
				isDelete := true
				for _, newEntry := range newConf {
					if logEntry.Path == newEntry.Path && logEntry.Topic == newEntry.Topic {
						isDelete = false
						break
					}
				}
				if isDelete {
					mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
					tskMgr.tskMap[mk].cancelFunc()
				}
			}

			fmt.Println("新的配置来了", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数，向外暴露tskMgr的NewConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}

// 一个函数，向tskMgr的NewConfChan推送数数据
//func PushNewConf(newConf []*etcd.LogEntry) {
//	tskMgr.newConfChan <- newConf
//}
