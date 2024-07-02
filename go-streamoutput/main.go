package main

import (
	"context"
	"demo/go-streamoutput/logic"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", DemoServer)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}

func DemoServer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	buildChat := logic.BuildChat{
		Ctx:                ctx,
		HttpResponseWriter: w,
		HttpRequest:        r,
	}
	buildChat.BuildChat(r)
}

//普通响应示例：
/*
HTTP/1.1 200 OK
Content-Type: application/json
Connection: close

{"message": "Data retrieved successfully", "data": {...}}
*/
//流式输出示例：
/*
data: {"message": "New message", "timestamp": "2024-06-13T12:34:56Z"}
\n\n
*/
