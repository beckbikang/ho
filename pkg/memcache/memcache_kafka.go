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

	kafka.SendTokafka(key, sendData)
	res.Response = mc.RespStored
	return nil
}
