package main

import (
	"godemo/go-daemon/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	now := time.Now()
	hour, minute, second := now.Clock()

	// 计算到今天23:59的持续时间
	duration := time.Duration((23-hour)*60*60+(60-minute)*60+(59-second)) * time.Second

	// 创建定时器
	ticker := time.NewTimer(duration)
	done := make(chan bool, 1) //结束任务信号器

	//处理退出信号，优雅的关闭定时器
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		log.Println("收到停止信号！")
		ticker.Stop()
		done <- true
	}()

	//定时任务执行
	log.Println("The task is executing......")
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				service.RunTask()
			}
		}
	}()

	//阻塞主线程，让守护进程持续运行
	select {}
}
