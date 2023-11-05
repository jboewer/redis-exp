package redis

import (
	"github.com/jboewer/redis/pkg/resp"
	"net"
)

func cmdPing(conn net.Conn, cmd command) {
	var answer string
	if len(cmd.args) > 0 {
		answer = cmd.args[0]
	} else {
		answer = "PONG"
	}

	s, err := resp.NewSimpleString(answer)
	if err != nil {
		return
	}
	_, err = conn.Write(s.Bytes())
	return
}
