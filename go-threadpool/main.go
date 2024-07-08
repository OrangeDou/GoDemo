package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, tasks chan int) {
	for task := range tasks {
		//do task logic
		fmt.Printf("do task %d\n", task)
		time.Sleep(time.Second)
		wg.Done()
	}
}

func distributeTask(taskCount int, workerCount int) {
	var wg sync.WaitGroup

	task := make(chan int, taskCount)

	for i := 0; i < workerCount; i++ {
		go worker(&wg, task)
	}
	for i := 0; i < taskCount; i++ {
		wg.Add(1)
		task <- 1
	}
	// 关闭任务队列
	close(task)
	wg.Wait()
}

func main() {
	const taskCount = 50
	const workerCount = 5

	distributeTask(taskCount, workerCount)
}
