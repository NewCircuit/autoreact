package main

import (
	. "github.com/Floor-Gang/suggestions/internal"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := GetConfig("./config.yml")
	Start(config)
	keepAlive()
}

func keepAlive() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
