package main

import (
	"fmt"
	"sync"
)

type clock interface {
	add()
	read() int
}

type clocker struct {
	val  int
	lock sync.Mutex
}

func (c *clocker) add() {
	c.lock.Lock()
	c.val++
	c.lock.Unlock()
}
func (c *clocker) read() int {
	return c.val

}

func main() {
	clocker := &clocker{}
	var wg sync.WaitGroup
	isInterface(clocker)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clocker.add()
		}()
	}
	wg.Wait()
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(clocker.val)
		}()
	}
	wg.Wait()
}

func isInterface(clock clock) bool {
	clock.add()
	return true
}
