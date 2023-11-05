package redis

import (
	"fmt"
	"github.com/jboewer/redis/pkg/resp"
	"net"
)

type Server struct {
}

func (s *Server) ListenAndServe(addr string) error {
	l, err := net.Listen("tcp", addr)

	if err != nil {
		return fmt.Errorf("failed to bind to addr %s\n", addr)
	}

	return s.Serve(l)
}

func (s *Server) Serve(l net.Listener) error {
	defer l.Close()
	for {
		conn, err := l.Accept()

		if err != nil {
			return fmt.Errorf("failed accepting connection %s\n", err.Error())
		}

		go func() {
			err := s.handleConnection(conn)
			if err != nil {
				fmt.Printf("failed handling connection %s\n", err.Error())
			}
		}()
	}
}

func (s *Server) handleConnection(conn net.Conn) error {
	defer conn.Close()

	rd := resp.NewReader(conn)
	for {
		v, err := rd.ReadValue()
		if err != nil {
			return fmt.Errorf("failed reading value %s\n", err.Error())
		}

		fmt.Printf("Recieved %s\n", v.Type())
	}
}
