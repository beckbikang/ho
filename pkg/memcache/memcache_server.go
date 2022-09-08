package memcache

import (
	"context"
	"ho/pkg/global"
	"strconv"
	"sync"

	mc "github.com/rpcxio/gomemcached"
)

var memStore sync.Map

func GetMemcacheServer() *mc.Server {
	addr := global.GCONFIG.GetString("main.ServerIp") + ":" + strconv.Itoa(global.GCONFIG.GetInt("main.ServerPort"))
	mcServer := mc.NewServer(addr)
	mcServer.RegisterFunc("set", defaultSet)
	return mcServer
}

func defaultSet(ctx context.Context, req *mc.Request, res *mc.Response) error {
	key := req.Key
	value := req.Data
	memStore.Store(key, value)
	res.Response = mc.RespStored
	return nil
}
