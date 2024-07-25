package main

import (
	"fmt"
	"rabit/service"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//请求写入rabbitMQ
	wg.Add(1)
	go service.SendRequestsToQueue(10000, &wg)

	wg.Wait()
	fmt.Println("Sent all requests")

	//进行消费
	service.Consumer()
}
