package main

import (
	"fmt"
	"github.pie.apple.com/privatecloud/dcdr/client"
	"github.pie.apple.com/privatecloud/dcdr/config"
	"net/http"
	"os"
	"simple-go-http-rest/api"
)

func main() {
	http.HandleFunc("/", index)

	http.HandleFunc("/api/booleanEcho", api.BooleanEchoHandleFunc)
	http.HandleFunc("/api/percentEcho", api.PercentEchoHandleFunc)

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
			OutputPath: "/etc/dcdr/decider.json",
		},
	}

	client, err := client.New(config)

	if err != nil {
		panic(err)
	}

	// example-feature would be false as the id is not set, but the flag has id's associated with it.
	if client.IsAvailable("2pac/newish", "") {

		fmt.Fprintf(w, "Daniel's exciting and looney world. flag enabled")
	} else {
		fmt.Fprintf(w, "Daniel's sane and boring world. flag disabled")
	}

}
