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

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:topic`
}

// 初始化etcd的函数

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

// 从etcd中根据key 获取配置项

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