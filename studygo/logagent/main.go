package main

import (
	"fmt"
	"go_git/Go_learn/studygo/logagent/conf"
	"go_git/Go_learn/studygo/logagent/etcd"
	"go_git/Go_learn/studygo/logagent/kafka"
	"go_git/Go_learn/studygo/logagent/taillog"
	"go_git/Go_learn/studygo/logagent/utils"
	"time"
	"sync"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2.发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}

func main() {
	// 0. 加载配置文件
	// cfg, err := ini.Load("./conf/config.ini")
	// fmt.Println(cfg.Section("kafka").Key("address").String())
	// fmt.Println(cfg.Section("kafka").Key("topic").String())
	// fmt.Println(cfg.Section("taillog").Key("path").String())
	//return
	err := ini.MapTo(cfg, "./conf/config.ini")
	fmt.Println(cfg.KafkaConf.Address)
	fmt.Println(cfg.KafkaConf.Topic)
	fmt.Println(cfg.TaillogConf.FileName)
	//return
	if err != nil {
		fmt.Printf("load ini failed ,err:%v\n", err)
		return
	}

	//1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}，cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")

	// 2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")

	// 2.0 为了实现每个logagent都拉取自己独有的配置，所以要以自己的IP地址作为区分
	ipStr,err := utils.GetOutboundIP()
	if err !=nil{
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.key,ipStr)
	// 2.1从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v", err)
		return
	}
	fmt.Printf("etcd.GetConf success. %v\n", logEntryConf)
	// 2.2派一个哨兵监视日志收集项的变化（watch），实现热加载配置

	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v", index, value)
	}
	// 3.收集日志发往kafka
	taillog.Init(logEntryConf)	
	newConfChan := taillog.NewConfChan()	// 从taillog中获取对外暴露的通道
	var wg  sync.WatiGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey,NewConfChan) // 哨兵发现最新的配置信息会通知上面的通道
	wg.Wait()
	
}
