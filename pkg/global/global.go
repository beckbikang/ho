package global

import (
	"log"

	"ho/internal/config"
	"ho/pkg/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GF *config.GlobalConf
var LOGGER *logger.Logger

// 全局配置
var GCONFIG *viper.Viper

func InitConfig(configPath, configFileName string) {
	GF = new(config.GlobalConf)
	GCONFIG = viper.New()

	log.Printf("%s/%s", configPath, configFileName)

	GCONFIG.AddConfigPath(configPath)
	GCONFIG.SetConfigName(configFileName)
	GCONFIG.SetConfigType("toml")

	if err := GCONFIG.ReadInConfig(); err != nil {
		log.Printf("error :%v\n", err)
		panic(-1)
	}

	err := GCONFIG.Unmarshal(GF)

	if err != nil {
		log.Printf("error :%v\n", err)
		panic(-1)
	}

	//监听配置变化
	GCONFIG.WatchConfig()
	GCONFIG.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	log.Println("start init")

	//init server
	InitComponents()

}

func InitComponents() {
	//init log
	lfg := new(logger.LogConfig)
	lfg.Filename = GCONFIG.GetString("main.MainLogPath")
	LOGGER = logger.NewLogger(lfg)
}
