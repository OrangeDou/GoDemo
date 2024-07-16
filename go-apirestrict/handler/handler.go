package handler

import (
	"demo/go-apirestrict/redis"
	"encoding/json"
	"log"
	"net/http"
)

type Msg struct {
	Message string `json:"message"`
}

// 测试api
func AipTest(w http.ResponseWriter, r *http.Request) {
	//判断redis计数器是否已满，未满则允许访问,否则拒绝访问
	result, err := redis.IsFull()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	if !result {
		w.WriteHeader(http.StatusForbidden)
	} else {
		//接口响应部分,访问成功后将访问消息写入redis，并设置计数器+1操作
		r.Header.Set("Content-Type", "application/json")
		msg := Msg{Message: "Hello, test!"}
		msgByte, err := json.Marshal(msg)
		if err != nil {
			log.Print(err.Error())
		}
		redis.ConnectRedis()
		w.Write(msgByte)
	}
}
