package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

// 1. 实现三节点选举
// 2. 改造代码成分布工选举代码，加入RPC调用
// 3. 演示完整代码， 自动选主，日志复制

// 定义3节点常量
const RaftCount = 3

// 声明 leader对象
type Leader struct {
	//任期
	Term int
	//Leader编号
	LeaderId int
}

// 声明
type Raft struct {
	// 锁
	mu sync.Mutex
	// 节点编号
	me int
	// 当前任期
	currentTerm int
	// 为哪个节点投票
	votedFor int
	// 3个状态
	// 0 follower  1 canidate  2 leader
	state int
	// 发送最后一条数据的时间
	lastMessageTime int64
	// 当前节点的Leader
	currentLeader int
	// 节点间发信息的通道
	message chan bool
	// 选举的通道
	electCh chan bool
	// 心跳信号的通道
	heartBeat chan bool
	// 返回心跳信号的通道
	heartbeatRe chan bool
	// 超时时间
	timeout int
}

// 0 没有上任  -1 没胡编号
var leader = Leader{0, -1}

func main() {
	// 过程，3个节点，最初都是follower
	// 若有candidate状态， 进行投票拉票
	// 会产生leader

	for i := 0; i < RaftCount; i++ {
		// 创建3个节点
		Make(i)
	}

	// 加入服务端注册
	rpc.Register(new(Raft))
	rpc.HandleHTTP()
	// 监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	for {

	}
}

func Make(me int) *Raft {
	rf := &Raft{
		me: me,
		//  -1 代表谁都没投，此时节点刚创建
		votedFor: -1,
		// 0 follower
		state:           0,
		lastMessageTime: 0,
		currentLeader:   -1,
		message:         make(chan bool),
		electCh:         make(chan bool),
		heartBeat:       make(chan bool),
		heartbeatRe:     make(chan bool),
		timeout:         0,
	}
	rf.setTerm(0)
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	//选举的协程
	go rf.election()

	// 心跳检测的协程
	go rf.sendLeaderHeartBeat()

	return rf
}

func (raft *Raft) setTerm(term int) {
	raft.currentTerm = term
}

// 获取当前时间，发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (rf *Raft) election() {
	//设置标记， 判断是否选出了Leader
	var result bool
	result = false
	for {
		//设置超时， 150到300的随机数
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()
		select {
		//延迟等待1毫秒
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点状态为：", rf.state)
		}
		for !result {
			//选主逻辑
			result = rf.election_one_round(&leader)
		}
	}
}

// 随机值
func randRange(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 实现选主的逻辑
func (rf *Raft) election_one_round(leader *Leader) bool {
	//定义超时
	var timeout int64
	timeout = 100
	// 投票数量
	var vote int
	// 定义是否开始心跳信号的产生
	var triggerHeartBeat bool
	var success bool
	// 时间
	last := millisecond()

	//给当前节点变成 candidate
	rf.mu.Lock()
	// 修改状态
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start electing leader...")
	for {
		// 遍历所有节点拉选票
		for i := 0; i < RaftCount; i++ {
			if i != rf.me {
				// 拉选票
				go func() {
					if leader.LeaderId < 0 {
						//设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量
		vote = 0
		// 遍历节点
		for i := 0; i < RaftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-rf.electCh:
				if ok {
					// 投票数量加 1
					vote = vote + 1
					success = vote > RaftCount/2
					if success && !triggerHeartBeat {
						// 变化为 Leader,选主成功
						// 开始触发心跳信号检测
						rf.mu.Lock()
						rf.becomeLeader()
						triggerHeartBeat = true
						rf.mu.Unlock()
						// 由 leader 向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为Leader")
						fmt.Println("leader开始发送心跳信号了")
					}
				}
			}
		}
		//做最后的检验工作
		// 若不超时，且票数大于一半，则选举成攻，break
		if timeout+last < millisecond() || (vote > RaftCount/2 || rf.currentLeader > 0) {
			break
		} else {
			//等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

// 修改状态为 candidate
func (rf *Raft) becomeCandidate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.votedFor = rf.me
	rf.currentLeader = -1
}

// 修改状态为 Leader
func (rf *Raft) becomeLeader() {
	rf.state = 2
	rf.currentLeader = rf.me
}

// leader 节点发送心跳信号
// 顺便完成数据同步
// 看小滴挂没挂
func (rf *Raft) sendLeaderHeartBeat() {
	//
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntries()
		}
	}
}

// 用于返回给leader的确认信号
func (rf *Raft) sendAppendEntries() {
	// 是主就别跑下面代码
	if rf.currentLeader == rf.me {
		//此时是leader
		//记录 确认信号的节点个数
		var success_count = 0
		// 设置确认信号
		for i := 0; i < RaftCount; i++ {
			if i != rf.me {
				go func() {
					//rf.heartbeatRe <- true
					// 这里实际上相当于客户端
					rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8000")
					if err != nil {
						log.Fatal("dialing:", err)
					}
					// 接收服务器返回的信息
					var ok = false
					err = rp.Call("Raft.Communication", Param{"hello"}, &ok)
					if err != nil {
						log.Fatal("error:", err)
					}
					if ok {
						rf.heartbeatRe <- true
					}
				}()
			}
		}
		//计算返回确认信号个数
		for i := 0; i < RaftCount; i++ {
			select {
			case ok := <-rf.heartbeatRe:
				if ok {
					success_count++
					if success_count > RaftCount/2 {
						fmt.Println("投票选举成功，心跳信号OK")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}

}

// 首字母大写， RPC规范
// 分布式通信
type Param struct {
	Msg string
}

// 通信方法
func (rf *Raft) Communication(p Param, a *bool) (err error) {
	fmt.Println(p.Msg)
	*a = true
	return nil
}
