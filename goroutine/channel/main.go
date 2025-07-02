package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "🐏"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "🐂"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	for {
		select {
		case <-c1:
			fmt.Println(c1)
		case <-c2:
			fmt.Println(c2)
		}
	}

}
