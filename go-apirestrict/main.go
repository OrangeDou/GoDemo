package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	wg      sync.WaitGroup
	taskNum int
)

func main() {
	taskNum = 100
	for i := 0; i < taskNum; i++ {
		wg.Add(1)
		go requestTask(&wg)
	}
	wg.Wait()
	fmt.Print("request has done!")
}

func requestTask(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get("http://localhost:8080/redis")
	if err != nil {
		log.Print(err)
	}

}
