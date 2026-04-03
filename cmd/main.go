package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func Request(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("END"))
}

func main() {

	limiter := NewFixedWindowRateLimiter(5, time.Second*5)
	http.HandleFunc("/", RateLimitMiddleware(limiter, Request))

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RateLimitMiddleware(limiter *FixedWindowRateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid IP", http.StatusInternalServerError)
			return
		}

		if limiter.Allow() {
			next(w, r)
		} else {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}
	}
}
