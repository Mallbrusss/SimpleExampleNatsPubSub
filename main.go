package main

import (
	"checkNats/consumer"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	consumer.Consumer()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	consumer.Disconnect()
	log.Println("Stop program")
}
