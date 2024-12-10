package conf

type Confs struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

type KafkaConf struct {
	Addr     string `ini:"addr"`
	ChanSize int    `ini:"chan_size""`
}

type EtcdConf struct {
	Addr string `ini:"addr"`
	Key  string `ini:"key"`
}
