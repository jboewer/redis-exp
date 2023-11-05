package redis

import (
	"errors"
	"fmt"
	"github.com/jboewer/redis/pkg/resp"
	"io"
	"net"
)

type command struct {
	name string
	args []string
}

type handleFunc func(conn net.Conn, cmd command)

type serveMux struct {
	handlers map[string]handleFunc
}

func (m *serveMux) Handle(conn net.Conn, cmd command) {
	if h, ok := m.handlers[cmd.name]; ok {
		h(conn, cmd)
		return
	}
	respErr, _ := resp.NewSimpleError(fmt.Sprintf("ERR unknown command '%s'", cmd.name))
	conn.Write(respErr.Bytes())
	fmt.Printf("unknown command %s\n", cmd.name)
}

type Server struct {
	mux *serveMux
}

func NewServer() *Server {

	mux := &serveMux{
		handlers: make(map[string]handleFunc),
	}

	mux.handlers["ping"] = cmdPing

	return &Server{
		mux: mux,
	}
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
			if err != nil && !errors.Is(err, io.EOF) {
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
			return fmt.Errorf("failed reading value %w", err)
		}

		// command must be an array
		vArr, ok := v.(resp.Array)
		if !ok {
			return fmt.Errorf("expected array, got %v\n", v.Type())
		}

		cmd, err := parseCommand(vArr)
		if err != nil {
			return fmt.Errorf("failed parsing command %s\n", err.Error())
		}

		s.mux.Handle(conn, cmd)
	}
}

func parseCommand(v resp.Array) (command, error) {
	var cmd command

	// All values must be bulk strings
	for _, v := range v.Values {
		bs, ok := v.(resp.BulkString)
		if !ok {
			return cmd, fmt.Errorf("expected bulk string, got %v\n", v.Type())
		}

		cmd.args = append(cmd.args, bs.Value)
	}

	cmd.name = cmd.args[0]

	return cmd, nil
}
