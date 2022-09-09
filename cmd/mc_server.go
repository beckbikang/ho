package cmd

import (
	"ho/pkg/global"
	"ho/pkg/memcache"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var mcServer = &cobra.Command{
	Use:   "mc-server",
	Short: "配置文件",
	Long:  "配置文件地址",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("mc to kafka")
		log.Println("configFilePath:" + configFilePath + " configFileName:" + configFileName)
		//init config
		global.InitConfig(mcConfigFilePath, mcConfigFileName)

		//init server
		mcServer := memcache.GetMemcacheServer()
		mcServer.RegisterFunc("set", memcache.DefaultSet)
		mcServer.RegisterFunc("get", memcache.DefaultGet)
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
				return
			}
		}
	},
}

var mcConfigFilePath string
var mcConfigFileName string

func init() {
	mcServer.Flags().StringVarP(&mcConfigFilePath, "path", "p", "", `配置文件路径`)
	mcServer.Flags().StringVarP(&mcConfigFileName, "filename", "f", "", `配置文件名称`)
}
