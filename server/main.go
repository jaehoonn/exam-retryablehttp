package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			log.Println("hello(GET)")
			w.Write([]byte("(GET) Hello World"))
		} else if req.Method == "PUT" {
			time.Sleep(5 * time.Second)
			log.Println("hello(PUT)")
			w.Write([]byte("(PUT) Hello World"))
		}

	})

	http.ListenAndServe(":5000", nil)
}
