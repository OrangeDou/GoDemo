package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const jsonString = `{"name": "John Doe", "age": 30, "is_active": true}`

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

func main() {
	var user User
	user.Name = "OrangeDou"
	user.Age = 25
	user.IsActive = true
	value, err := json.Marshal(user)
	if err != nil {
		return
	}
	http.HandleFunc("/getUserInfo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", value)

	})

	http.ListenAndServe(":8080", nil)
}

// 模拟高并发访问api
// func requestApi(url string) (times int) {

// }
