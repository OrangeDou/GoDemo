package service

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/streadway/amqp"
)

func sendRequest(ch *amqp.Channel, i int) {
	body, err := http.Get("http://localhost:8081/rabbit")
	if err != nil {
		return
	}
	respBody, err := io.ReadAll(body.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	err = ch.Publish(
		"",              // exchange
		"request_queue", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         respBody,
		})
	if err != nil {
		fmt.Println("Failed to publish message:", err)
	} else {
		fmt.Printf("Sent request: %s\n", body)
	}
}

func SendRequestsToQueue(requests int, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"request_queue", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		panic(err)
	}

	for i := 0; i < requests; i++ {
		sendRequest(ch, i)
	}
}
