package service

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Date struct {
	Response string `json:"res"`
}

var (
	wg *sync.WaitGroup
)

func Ask() error {
	data := Date{Response: "Hello haha!"}
	http.HandleFunc("/redis", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, err := json.Marshal(data)
		if err != nil {
			// 如果发生错误，返回 HTTP 500 内部服务器错误
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}
