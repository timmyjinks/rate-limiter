package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func Request(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("END"))

}

func main() {

	limiter := NewIPRateLimiter()
	http.HandleFunc("/", RateLimitMiddleware(limiter, Request))

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RateLimitMiddleware(limiter *IPRateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid IP", http.StatusInternalServerError)
			return
		}

		if limiter.Allow(ip) {
			next(w, r)
		} else {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}
	}
}
