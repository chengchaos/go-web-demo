package main

import (
	"fmt"
	"net/http"

	"github.com/chengchaos/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello %s\n", p.ByName("name"))
}

func root(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "<h1>It works!</h1><br /><a href='/hello/chengchao'>chengchao</a>")
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.GET("/", root)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	fmt.Println("server is startting ...")
	server.ListenAndServe()
	fmt.Println("server is stoped!")
}
