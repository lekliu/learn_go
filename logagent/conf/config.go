package conf

type AppConfig struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type EtcdConf struct {
	Address         string `ini:"address"`
	Timeout         int    `ini:"timeout"`
	Collect_log_key string `ini:"collect_log_key"`
}

// ---------------unUsed ↓ ----------------

type TaillogConf struct {
	Filename string `ini:"filename"`
}
