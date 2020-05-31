package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
	// --- unsed ---
	TaillogConf `ini:"taillog"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:chan_max_size`
	// --- unsed ---
	Topic string `ini:"topic"`
}
type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}

// --- unsed ---
type TaillogConf struct {
	FileName string `ini:"filename"`
}
