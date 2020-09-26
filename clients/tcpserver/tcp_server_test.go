package tcpServer

import (
	"bufio"
	"github.com/tomiok/fuego-cache/internal"
	"net"
	"testing"
	"time"
)

func newTestServer() *server {
	return New("localhost:11111")
}

func Test_TCPServer(t *testing.T) {
	s := newTestServer()

	s.onNewMessage = func(c *Client, message string) {
		_ = c.Send(message)
	}

	go s.Listen()

	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:11111")

	if err != nil {
		t.Fail()
	}

	if conn == nil {
		t.Fail()
		return
	}

	_, err = conn.Write([]byte("hello dude. \n"))

	if err != nil {
		t.Fail()
	}

	time.Sleep(100 * time.Millisecond)
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')

	if line != "hello dude. \n" {
		t.Fail()
	}

	internal.OnCloseError(conn.Close)
}
