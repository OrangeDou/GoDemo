package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type BuildChat struct {
	Ctx                context.Context
	HttpResponseWriter http.ResponseWriter
	HttpRequest        *http.Request
}

type BuildChatPush struct {
	Code int64              `json:"code"`
	Msg  string             `json:"msg"`
	Time time.Time          `json:"diff_time,omitempty"`
	Data *BuildChatPushData `json:"data,omitempty"`
}

type BuildChatPushData struct {
	Content   string `json:"content"`    // 回复内容
	IsSuccess int64  `json:"is_success"` // 是否成功
}

func (l *BuildChat) Push(reply *BuildChatPush) error {
	//接口变量类型断言，判断当前是否为Flusher类型
	//http.Flusher接口允许强制HTTP响应流发送立即刷新，这对于流式输出非常有用。
	flusher, ok := l.HttpResponseWriter.(http.Flusher)
	if !ok {
		return errors.New("Streaming unsupported")
	}
	message, _ := json.Marshal(reply)
	_, err := fmt.Fprintf(l.HttpResponseWriter, "data: %s\n\n", string(message))
	if err != nil {
		return err
	}

	flusher.Flush() //强制HTTP响应流立即发送，确保客户端能够及时接收到数据
	return nil
}

func (l *BuildChat) BuildChat(req *http.Request) error {
	l.Push(&BuildChatPush{
		Code: 100,
		Msg:  "This is the stream word",
		Time: time.Now(),
	})

	return nil
}
