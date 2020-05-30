package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
	// --- unsed ---
	TaillogConf `ini:"taillog"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	// --- unsed ---
	Topic string `ini:"topic"`
}
type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
}

// --- unsed ---
type TaillogConf struct {
	FileName string `ini:"filename"`
}