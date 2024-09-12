package conf

type LogTransferConfig struct {
	KafkaConf `ini:"kafka"`
	EsConf    `ini:"es"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EsConf struct {
	Address    string `ini:"address"`
	ChanSize   int    `ini:"chan_size"`
	ChanNumber int    `ini:"chan_number"`
}
