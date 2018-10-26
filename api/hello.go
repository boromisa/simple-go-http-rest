package api

import (
	"fmt"
	"github.pie.apple.com/privatecloud/dcdr/client"
	"github.pie.apple.com/privatecloud/dcdr/config"
	"net/http"
	"strconv"
)

// PercentEchoHandleFunc to be used as http.HandleFunc for Hello API
func PercentEchoHandleFunc(w http.ResponseWriter, r *http.Request) {

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
	var output string

	if client.IsAvailableForPercentile("2pac/percent", message) {
		output += fmt.Sprintf("The whitelisted unicorn user ID: %s percent ENABLED\n", message)
	} else {
		output += fmt.Sprintf("The black sheep user ID: %s percent disabled\n", message)
	}

	for i := 0; i < 100; i++ {

		if client.IsAvailableForPercentile("2pac/percent", message+strconv.Itoa(i)) {
			output += fmt.Sprintf("Loop number : %s with the user ID: %s example-feature ENABLED\n", strconv.Itoa(i), message+strconv.Itoa(i))
		} else {
			output += fmt.Sprintf("Loop number : %s with the user ID: %s example-feature disabled\n", strconv.Itoa(i), message+strconv.Itoa(i))
		}

	}

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, output)
}
