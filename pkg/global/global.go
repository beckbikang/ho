package global

import (
	"log"

	"ho/internal/config"
	"ho/pkg/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const GLOBAL_VIP_SLITE = "."

var GF *config.GlobalConf
var LOGGER *zap.Logger

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
	initLog()
}

func initLog() {
	//init main log file
	lfg := new(logger.LogConfig)
	logger.WithMultiFile(lfg)
	lfg.Filename = GCONFIG.GetString("main.mainLogPath")
	lfg.LogMod = GF.MainCfg.MainLogModel
	LOGGER = logger.NewLogger(lfg).GetZlog()
}
