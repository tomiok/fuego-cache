package stdioClient

import (
	"bufio"
	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
	"os"
	"os/signal"
)

type callback func(string) string

type StdClient struct {
	OnMessage callback
}

func NewStdClient() *StdClient {
	return &StdClient{}
}

func (s *StdClient) OnNewMessage(callback callback) {
	s.OnMessage = callback
}

func (s *StdClient) StandardInputCache() {
	quit := make(chan os.Signal, 1)
	go func() {
		logs.StdInfo("start with fuego here... (set 1 1)")
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			logs.StdInfo(s.OnMessage(text))
		}
	}()

	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		logs.StdInfo("exiting...")
	}
}

func (s *StdClient) PrintBanner() {
	internal.PrintBanner()
}
