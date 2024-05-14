package main

import (
	"net/http"
	"time"
)

func main() {
	server := &http.Server{Addr: ":3000"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World"))
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
