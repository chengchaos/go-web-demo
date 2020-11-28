package main

import "net/http"
import "fmt"

type HelloHandler struct {

}

type WorldHandler struct {

}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world!")
}

func main() {

	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr : "127.0.0.1:8080",

	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}