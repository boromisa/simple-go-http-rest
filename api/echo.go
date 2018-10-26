package api

import (
	"fmt"
	"github.pie.apple.com/privatecloud/dcdr/client"
	"github.pie.apple.com/privatecloud/dcdr/config"

	"net/http"
)

// BooleanEchoHandleFunc to be used as http.HandleFunc for ECHO API
func BooleanEchoHandleFunc(w http.ResponseWriter, r *http.Request) {

	config := &config.Config{
		Watcher: config.Watcher{
			OutputPath: "/etc/dcdr/decider.json",
		},
	}

	client, err := client.New(config)

	if err != nil {
		panic(err)
	}

	message := r.URL.Query()["message"][0]

	// example-feature would be false
	if client.IsAvailable("2pac/newish", message) {
		message = message + " example-feature enabled"
	} else {
		message = message + " example-feature disabled"
	}

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
