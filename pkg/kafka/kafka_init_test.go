package kafka

import (
	"ho/pkg/global"
	"log"
	"sync"
	"testing"
)

func TestInitKafka(t *testing.T) {
	configPath := "../../configs"
	configFileName := "config-dev.toml"

	global.InitConfig(configPath, configFileName)

	kafkaMap := global.GCONFIG.GetStringMap(KAFKA_KEY_NAME)
	log.Printf("kafkaMap:+%v", kafkaMap)

	brockersList := global.GCONFIG.GetStringSlice("kafkas.test_mc.brokers")
	log.Printf("######Brokers %v", brockersList)

	err := InitKafka()
	if err != nil {
		log.Fatal(err)
	}

	//init data
	proudcer1 := GetProducer("test_mc")
	defer proudcer1.Close()

	//sync send data
	topicName := global.GCONFIG.GetString("kafkas.test_mc.topic")
	log.Printf("send to topic %s", topicName)
	proudcer1.SyncSend(topicName, "{\"ho\":1}")

	var wg sync.WaitGroup
	wg.Add(1)
	proudcer1.Send(topicName, "{\"ho\":2}", func(meta *RecordMetadata, err error) {
		wg.Done()
	})
	wg.Wait()

}
