package main

import (
	"fmt"
	"go_git/Go_learn/studygo/logagent/conf"
	"go_git/Go_learn/studygo/logagent/etcd"
	"go_git/Go_learn/studygo/logagent/kafka"
	"go_git/Go_learn/studygo/logagent/taillog"
	"time"

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
	err = kafka.Init([]string{cfg.KafkaConf.Address})
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

	// 2.1从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf("/xx")
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v", err)
		return
	}
	fmt.Printf("etcd.GetConf success. %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v", index, value)
	}
	// 2.2派一个哨兵监视日志收集项的变化（watch），实现热加载配置

	// 3.打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v", err)
		return
	}
	fmt.Println("init taillog success.")
	run()

}
