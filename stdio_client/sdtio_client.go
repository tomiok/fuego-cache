package stdioClient

import (
	"bufio"
	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
	"os"
	"os/signal"
)

func StandardInputCache() {
	quit := make(chan os.Signal, 1)
	go func() {
		for {
			logs.StdInfo("start with fuego here... (set 1 1)")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			logs.Info(text)
		}
	}()

	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		logs.StdInfo("exiting...")
	}
}

func PrintBanner() {
	internal.PrintBanner()
}
