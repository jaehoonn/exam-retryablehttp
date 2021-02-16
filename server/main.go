package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		log.Println("hello")
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":5000", nil)
}
