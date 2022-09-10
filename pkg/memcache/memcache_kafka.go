package memcache

import (
	"context"
	"ho/pkg/global"

	mc "github.com/rpcxio/gomemcached"
	"go.uber.org/zap"
)

func DefaultSet2(ctx context.Context, req *mc.Request, res *mc.Response) error {
	key := req.Key
	value := req.Data
	global.LOGGER.Info("mc-set", zap.String(key, string(value)))
	memStore.Store(key, value)
	res.Response = mc.RespStored
	return nil
}
