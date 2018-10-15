package main

import (
	"github.com/boromisa/dcdr/client"
	"github.com/boromisa/dcdr/config"
	"simple-go-http-rest/api"
	"fmt"
	"net/http"
	"os"


)

func main() {



	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/hello", api.HelloHandleFunc)
	

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	config := &config.Config{
	Watcher: config.Watcher{
		OutputPath: "/etc/dcdr/dcdr.json",
	},
	}

	client, err := client.New(config)

	if err != nil {
		panic(err)
	}

	// example-feature would be false
	if client.IsAvailable("example-feature") {
		fmt.Fprintf(w, "Welcome to Cloud Native Go (Update). flag enabled")
	} else {
		fmt.Fprintf(w, "Welcome to Cloud Native Go (Update). flag disabled")
	}

}
