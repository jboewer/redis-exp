package redis

import (
	"github.com/jboewer/redis/pkg/resp"
	"net"
)

func cmdPing(conn net.Conn, cmd command) {
	s, err := resp.NewSimpleString("PONG")
	if err != nil {
		return
	}
	_, err = conn.Write(s.Bytes())
	return
}
