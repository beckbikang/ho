package redis

import (
	"ho/pkg/global"
	"ho/pkg/kafka"
	"log"
	"strconv"
	"strings"

	"github.com/tidwall/redcon"
	"go.uber.org/zap"
)

func InitRedisServer() error {
	addr := global.GCONFIG.GetString("main.serverIp") + ":" + strconv.Itoa(global.GCONFIG.GetInt("main.serverPort"))

	err := redcon.ListenAndServe(addr,
		func(conn redcon.Conn, cmd redcon.Command) {
			switch strings.ToLower(string(cmd.Args[0])) {
			default:
				conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
			case "ping":
				conn.WriteString("PONG")
			case "quit":
				conn.WriteString("OK")
				conn.Close()
			case "set":
				if len(cmd.Args) != 3 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				key := string(cmd.Args[1])
				value := string(cmd.Args[2])
				global.LOGGER.Info("send-to-kafka", zap.String(key, value))
				kafka.SendTokafka(key, value)
				conn.WriteString("OK")
			}
		},
		func(conn redcon.Conn) bool {
			return true
		},
		func(conn redcon.Conn, err error) {
			// This is called when the connection has been closed
			// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
