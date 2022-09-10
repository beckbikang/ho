package memcache

import (
	"ho/pkg/global"
	"strconv"
	"sync"

	mc "github.com/rpcxio/gomemcached"
)

var memStore sync.Map

func GetMemcacheServer() *mc.Server {
	addr := global.GCONFIG.GetString("main.serverIp") + ":" + strconv.Itoa(global.GCONFIG.GetInt("main.serverPort"))
	mcServer := mc.NewServer(addr)
	return mcServer
}
