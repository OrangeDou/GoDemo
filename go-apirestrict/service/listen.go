package main

import (
	"demo/go-apirestrict/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/redis", handler.AipTest)

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("HTTP server failed: %v", err)
		if err == http.ErrServerClosed {
			log.Println("Server closed gracefully")
		}
	}

}
