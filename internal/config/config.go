package config

type GlobalConf struct {
	Title   string   `toml:"title"`
	MainCfg MainConf `toml:"main"`
}

type MainConf struct {
	ServerIp    int `toml:"ServerIp"`
	ServerPort  int `toml:"ServerPort"`
	MainLogPath int `toml:"MainLogPath"`
}
