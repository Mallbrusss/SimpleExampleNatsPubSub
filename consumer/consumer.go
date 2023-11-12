package consumer

import (
	"fmt"
	"log"
	_ "time"

	_ "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

var natConn stan.Conn
var nutSubscription stan.Subscription

func Consumer() {
	natsURL := "nats://localhost:4222" // Замените на ваш URL NATS Streaming сервера
	clusterID := "nat1"                // Замените на ваш Cluster ID
	clientID := "nat_client"           // Замените на ваш Client ID
	channelName := "order_chanel"      // Замените на имя вашего канала

	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Ошибка при подключении к NATS Streaming: %v", err)
	}
	natConn = conn

	subscription, err := conn.Subscribe(channelName, func(msg *stan.Msg) {

		fmt.Printf("Получено сообщение: %s\n", string(msg.Data))

	})
	nutSubscription = subscription
	if err != nil {
		log.Fatalf("Ошибка при подписке на канал: %v", err)
	}
}

func Disconnect() {
	handleUnsubscribe()
	handleDisconnect()
}

func handleUnsubscribe() {
	err := nutSubscription.Close()
	if err != nil {
		return
	}
}

func handleDisconnect() {
	err := natConn.Close()
	if err != nil {
		return
	}
}
