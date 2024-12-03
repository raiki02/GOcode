package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	//TailfConf `ini:"tailf"`
	EtcdConf `ini:"etcd"`
}

type KafkaConf struct {
	//Topic   string `ini:"topic"'`
	Address string `ini:"address"'`
	MaxSize int    `ini:"chan_max_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}

type TailfConf struct {
	FileName string `ini:"filename"`
}
