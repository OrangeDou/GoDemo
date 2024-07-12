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

		//不设置以下属性的话，会由于请求不到https安全证书而拒绝访问，因为反向代理的链接时https百度，可以代理其他的资源
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
