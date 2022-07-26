package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dacapoday/httpproxy"
)

var build = "unknown"

func init() {
	fmt.Println("build:", build)
}

func main() {
	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	fmt.Printf("Listening on %v\n", addr)
	proxy := &httpproxy.Proxy{}
	http.ListenAndServe(addr, log(proxy))
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request Method:%v Host:%v \n", r.Method, r.URL.Host)
		h.ServeHTTP(w, r)
	})
}
