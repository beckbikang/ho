package kafka

import (
	"ho/pkg/global"
)

const KAFKA_KEY_NAME = "kafkas"

// parse global config to Options . then start kafka
func InitKafka() error {

	kafkaMap := global.GCONFIG.GetStringMap(KAFKA_KEY_NAME)
	global.LOGGER.Sugar().Infof("kafkaMap:+%v", kafkaMap)
	opts := make([]*Options, len(kafkaMap))

	i := 0
	for key, value := range kafkaMap {

		preConfigKey := KAFKA_KEY_NAME + global.GLOBAL_VIP_SLITE + key + global.GLOBAL_VIP_SLITE

		global.LOGGER.Sugar().Infof("kafkaconfig:%v,%v", key, value)
		opt := NewDefaultOptions()
		opt.Name = key

		global.LOGGER.Info("####Brokers:key" + preConfigKey + "brokers")
		global.LOGGER.Sugar().Infof("####Brokers:%v", global.GCONFIG.GetStringSlice(preConfigKey+"brokers"))
		opt.Addr = global.GCONFIG.GetStringSlice(preConfigKey + "brokers")

		opt.Consumer.Group = global.GCONFIG.GetString(preConfigKey + "group")
		opt.SASLEnable = global.GCONFIG.GetBool(preConfigKey + "sslEnable")
		opt.SASLUser = global.GCONFIG.GetString(preConfigKey + "user")
		opt.SASLPassword = global.GCONFIG.GetString(preConfigKey + "pswd")
		opt.NotNeedProducer = !global.GCONFIG.GetBool(preConfigKey + "producerOn")
		global.LOGGER.Sugar().Infof("opt[%s]:%v", key, *opt)
		opts[i] = opt
		i++
	}
	global.LOGGER.Sugar().Infof("opts:+%v", opts)
	global.LOGGER.Info("################start make kafka#######################")
	err := Init(opts)
	if err != nil {
		global.LOGGER.Sugar().Errorf("error:%v", err.Error())
		return err
	}
	return nil
}
