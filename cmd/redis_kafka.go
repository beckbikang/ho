package cmd

import (
	"ho/pkg/global"
	"ho/pkg/kafka"
	"ho/pkg/redis"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "运行服务",
	Long:  "运行服务",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var redisKafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "配置文件",
	Long:  "配置文件地址",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("mc to kafka")
		log.Println("redisConfigFilePath:" + redisConfigFilePath + " redisConfigFileName:" + redisConfigFileName)
		//init config
		global.InitConfig(redisConfigFilePath, redisConfigFileName)

		//init kafka
		kafka.InitKafka()

		//init server
		err := redis.InitRedisServer()
		if err != nil {
			panic(err)
		}

		//sign
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

		//wait
		for {
			select {
			case <-signChan: //stop sign
				global.LOGGER.Info("end")
				kafka.StopAll()
				return
			}
		}
	},
}

var redisConfigFilePath string
var redisConfigFileName string

func init() {
	redisCmd.AddCommand(redisKafkaCmd)
	redisKafkaCmd.Flags().StringVarP(&redisConfigFilePath, "path", "p", "", `配置文件路径`)
	redisKafkaCmd.Flags().StringVarP(&redisConfigFileName, "filename", "f", "", `配置文件名称`)
}
