package main

import (
	"log"
	"net/http"
	"os"

	web "github.com/s4mur4i/self-serve/pkg/http"
)

func main() {
	httpPort := os.Getenv("PORT")

	if httpPort == "" {
		httpPort = ":4444"
	} else {
		httpPort = ":" + httpPort
	}

	webService := web.New()

	log.Printf("Running on port %s\n", httpPort)
	// log.Printf("%s\n", config.Conf.Okta.ClientID)
	log.Fatal(http.ListenAndServe(httpPort, webService.Router))
}
