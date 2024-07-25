package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/rabbit", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-type", "application-json")
		msg := "success to get rabbit"
		json.Marshal(msg)
		w.Write([]byte(msg))
	})
	http.ListenAndServe(":8081", nil)
}
