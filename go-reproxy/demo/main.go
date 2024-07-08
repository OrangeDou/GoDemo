package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/reproxy", func(w http.ResponseWriter, r *http.Request) {
		target, _ := url.Parse("https://www.baidu.com")
		reProxy := httputil.NewSingleHostReverseProxy(target)
		reProxy.Director = func(r *http.Request) {
			r.URL.Host = target.Host
			r.URL.Path = target.Path
			r.URL.Scheme = target.Scheme
			r.Host = target.Host

		}
		reProxy.ServeHTTP(w, r)
	})

	log.Print(http.ListenAndServe(":8080", nil))
}
