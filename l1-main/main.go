package main

import "fmt"
import "net/http"

type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	handler := MyHandler{}
	server := http.Server {
		Addr : "127.0.0.1:8080",
		Handler : &handler,
	}
	server.ListenAndServe()
}