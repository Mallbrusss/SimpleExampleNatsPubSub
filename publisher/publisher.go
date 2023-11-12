package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	natsURL := "nats://localhost:4222"
	clusterID := "nat1"
	clientID := "nat_client2"
	channelName := "order_chanel"

	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Ошибка при подключении к NATS Streaming: %v", err)
	}
	defer func(conn stan.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	go func() {
		for {
			number := 1
			data := fmt.Sprintf("hello world %d", number+1)

			err = conn.Publish(channelName, []byte(data))
			if err != nil {
				log.Printf("Ошибка при публикации сообщения: %v", err)
			} else {
				fmt.Println("Отправлена JSON-строка:", string(data))
			}

			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			number := rand.Intn(201) - 100
			data := fmt.Sprintf("hello world %d", number)

			err = conn.Publish(channelName, []byte(data))
			if err != nil {
				log.Printf("Ошибка при публикации сообщения: %v", err)
			} else {
				fmt.Println("Отправлена JSON-строка:", string(data))
			}

			time.Sleep(1 * time.Second)
		}

	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("Завершение программы...")
}
