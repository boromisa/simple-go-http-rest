package api

import (
	"fmt"
	"github.com/boromisa/dcdr/client"
	"github.com/boromisa/dcdr/config"
	"net/http"
)

// EchoHandleFunc to be used as http.HandleFunc for ECHO API
func EchoHandleFunc(w http.ResponseWriter, r *http.Request) {

	config := &config.Config{
		Watcher: config.Watcher{
			OutputPath: "/etc/dcdr/dcdr.json",
		},
	}


	client, err := client.New(config)

	if err != nil {
		panic(err)
	}

	message := r.URL.Query()["message"][0]

	// example-feature would be false
	if client.IsAvailable("example-feature") {
		message = message + " example-feature enabled"
	} else {
		message = message + " example-feature disabled"
	}

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
