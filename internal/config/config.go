package config

type GlobalConf struct {
	Title    string                 `toml:"title"`
	MainCfg  MainConf               `toml:"main"`
	KafkaCfg map[string]KafkaConfig `toml:"kafkas"`
}

type MainConf struct {
	ServerIp        int  `toml:"serverIp"`
	ServerPort      int  `toml:"serverPort"`
	MainLogPath     int  `toml:"mainLogPath"`
	MainLogModel    int8 `toml:"mainLogModel"`
	ShowSaramaDebug bool `toml:"showSaramaDebug"`
}

type KafkaConfig struct {
	Brokers    []string `toml:"brokers"`
	Topic      string   `toml:"topic"`
	Group      string   `toml:"group"`
	SslEnable  bool     `toml:"sslEnable"`
	User       string   `toml:"user"`
	Pswd       string   `toml:"password"`
	ProducerOn bool     `toml:"producerOn"`
}
