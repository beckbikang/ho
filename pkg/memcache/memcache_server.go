package memcache

import (
	"context"
	"ho/pkg/global"

	mc "github.com/rpcxio/gomemcached"
	"go.uber.org/zap"
)

func DefaultSet(ctx context.Context, req *mc.Request, res *mc.Response) error {
	key := req.Key
	value := req.Data
	global.LOGGER.Info("mc-set", zap.String(key, string(value)))
	memStore.Store(key, value)
	res.Response = mc.RespStored
	return nil
}

func DefaultGet(ctx context.Context, req *mc.Request, res *mc.Response) error {
	for _, key := range req.Keys {
		value, _ := memStore.Load(key)
		res.Values = append(res.Values, mc.Value{key, "0", value.([]byte), ""})
		global.LOGGER.Info("mc-get", zap.String(key, string(value.([]byte))))
	}
	res.Response = mc.RespEnd

	return nil
}
