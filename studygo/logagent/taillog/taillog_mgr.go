package taillog

import (
	"fmt"
	"go_git/Go_learn/studygo/logagent/etcd"
	"time"
	//"go_git/Go_learn/studygo/logagent/taillog"
)

var tskMgr *tailLogMgr

// tailTask 管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

// Init ...
func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
	}

	for _, logEntry := range logEntryConf {
		//logEntry: *etcd.logEntry
		//logEntry.Path 要收集的日志文件路径
		// 初始化的时候齐了多少个tailtask 都要记下来，为了后续判断方便
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}
	go tskMgr.run()
}

// 监听自己的newConfChan，有了新的配置过来之后就对应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 1.配置新增
					tailObj := newTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			// 找出原来t.logEntry 有，但是newConf中没有的，要删掉
			for _, c1 := range t.logENtry { // 从原来的配置中一次拿出配置项
				isDelete := true
				for _, c2 := range newConf { // 从新的配置中逐一比较
					if c2.Path == c1.Path && c2.Topic == c1.Tipic {
						idDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对于的这个tailObj停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					// t.tskMap[mk] ==> tailObj
					t.tskMap[mk].cancelFunc()
				}
			}
			// 2.配置删除
			// 3.配置变更
			fmt.Println("新的配置起来了!", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 个函数，向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.NewConfChan
}
