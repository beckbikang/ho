package cmd

import (
	"ho/pkg/global"
	"ho/pkg/kafka"
	"ho/pkg/memcache"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var mcCmd = &cobra.Command{
	Use:   "mc",
	Short: "运行服务",
	Long:  "运行服务",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mcKafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "配置文件",
	Long:  "配置文件地址",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("mc to kafka")
		log.Println("configFilePath:" + configFilePath + " configFileName:" + configFileName)
		//init config
		global.InitConfig(configFilePath, configFileName)

		//init kafka
		kafka.InitKafka()

		//init server
		mcServer := memcache.GetMemcacheServer()
		mcServer.RegisterFunc("set", memcache.McSendToKafka)
		mcServer.Start()

		//sign
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

		//wait
		for {
			select {
			case <-signChan: //stop sign
				global.LOGGER.Info("end")
				mcServer.Stop()
				kafka.StopAll()
				return
			}
		}
	},
}

var configFilePath string
var configFileName string

func init() {
	mcCmd.AddCommand(mcKafkaCmd)
	mcKafkaCmd.Flags().StringVarP(&configFilePath, "path", "p", "", `配置文件路径`)
	mcKafkaCmd.Flags().StringVarP(&configFileName, "filename", "f", "", `配置文件名称`)
}
