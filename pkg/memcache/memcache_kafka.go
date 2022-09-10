package memcache

import (
	"context"
	"ho/pkg/global"
	"ho/pkg/kafka"

	mc "github.com/rpcxio/gomemcached"
	"go.uber.org/zap"
)

func McSendToKafka(ctx context.Context, req *mc.Request, res *mc.Response) error {
	key := req.Key
	value := req.Data

	sendData := string(value)
	global.LOGGER.Info("send-to-kafka", zap.String(key, sendData))

	kafkaProduce := kafka.GetProducer(key)
	if kafkaProduce != nil {
		kafkaProduce.Send(key, sendData, func(meta *kafka.RecordMetadata, err error) {
			if err != nil {
				global.LOGGER.Sugar().Errorf("error:%v", err)
			} else {
				global.LOGGER.Sugar().Info(meta)
			}
		})
	}

	res.Response = mc.RespStored
	return nil
}
