package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	TailfConf `ini:"tailf"`
}

type KafkaConf struct {
	Topic   string `ini:"topic"'`
	Address string `ini:"address"'`
}

type TailfConf struct {
	FileName string `ini:"filename"`
}
