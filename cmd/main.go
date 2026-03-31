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

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid IP", http.StatusInternalServerError)
			return
		}
	}
}
