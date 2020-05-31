package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

// LogEntry 需要收集的日子的配置信息
type LogEntry struct {
	Path  string `json:"path"` // 日志存放的路径
	Topic string `json:topic`  // 日志要发往Kafka中的哪个topic
}

// Init 初始化etcd的函数
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v", err)
		return
	}

	fmt.Println("connect to etcd success")

	//defer cli.Close()
	return
}

// GetConf 从etcd中根据key 获取配置项
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	//get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)

	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed,err:%v\n", err)
	}
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s:%s\n", ev, Key, ev, Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%v\n", err)
			return
		}
	}
}

// etcd watch
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	// 从通道尝试取值（监视信息）
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v \n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			// 通知taillog.tskMgr
			// 1.先判断操作的类型
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed,err:%v\n", err)
			}
			// 如果是删除操作，手动传递一个空的配置项

			fmt.Printf("get newconf,%v\n", newConf)
			newConfCh <- newConf
		}
	}
}
