package main


import "fmt"
import "net/http"
//import "github.com/chengchaos/go-web-demo/data"

func handler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "<h1>It works! %s</h1>", request.URL.Path[1:])
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello world")
}
func main() {
	
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr : "0.0.0.0:8080",
		Handler: mux,
	}
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)
	//fmt.Println("hello world")
	server.ListenAndServe()
}

