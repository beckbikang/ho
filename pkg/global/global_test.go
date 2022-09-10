package global

import (
	"log"
	"testing"
)

func TestInitConfig(t *testing.T) {
	configPath := "../../configs"
	configFileName := "config-dev.toml"

	InitConfig(configPath, configFileName)
	LOGGER.Info("test-online")
	log.Println(GCONFIG.GetString("kafkas.k1.brockers"))

	log.Printf("%+v\n", GCONFIG.GetStringMap("kafkas"))

}
