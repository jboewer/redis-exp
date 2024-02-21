package redis

import "net"

type Conn struct {
	conn net.Conn
}
