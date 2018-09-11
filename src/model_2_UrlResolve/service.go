package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/post/", post)
	mux.HandleFunc("/redirect/", redirect)
	mux.HandleFunc("/json/", sendJSON)

	service := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	fmt.Println("http://127.0.0.1:8080")
	service.ListenAndServe()
}
