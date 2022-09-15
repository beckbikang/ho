package cmd

import (
	"ho/pkg/global"
	"ho/pkg/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var kafkaToKafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "配置文件",
	Long:  "配置文件地址",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("mc to kafka")
		log.Println("kafkaToKafkaConfigFilePath:" + kafkaToKafkaConfigFilePath + " kafkaToKafkaConfigFileName:" + kafkaToKafkaConfigFileName)
		log.Println("from:" + fromTopic + " toTopic:" + toTopic)

		//init config
		global.InitConfig(kafkaToKafkaConfigFilePath, kafkaToKafkaConfigFileName)

		//init kafka
		kafka.InitKafka()

		//consume and produce
		fromKafkaConsumer := kafka.GetConsumer(fromTopic)

		listenObj := &kafka.ListenerObj{
			ListenFn: func(message kafka.ConsumerMessage, ack *kafka.Acknowledgment) {
				global.LOGGER.Sugar().Info("message:", message)
				global.LOGGER.Sugar().Info("ackinfo:", ack)
				kafka.SendTokafka(toTopic, message.Value)
			},
		}
		fromKafkaConsumer.AddListener(fromTopic, listenObj)
		go fromKafkaConsumer.Start()
		global.LOGGER.Info("starting running....")

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

var fromTopic string
var toTopic string
var kafkaToKafkaConfigFilePath string
var kafkaToKafkaConfigFileName string

func init() {
	//config
	kafkaToKafkaCmd.Flags().StringVarP(&kafkaToKafkaConfigFilePath, "path", "p", "", `配置文件路径`)
	kafkaToKafkaCmd.Flags().StringVarP(&kafkaToKafkaConfigFileName, "filename", "f", "", `配置文件名称`)
	//topic
	kafkaToKafkaCmd.Flags().StringVarP(&fromTopic, "from", "r", "", `来源topic`)
	kafkaToKafkaCmd.Flags().StringVarP(&toTopic, "to", "t", "", `目标topic`)
}
