// Package main implements a simple HTTP server that responds with a greeting message.
// 
// The server listens on port 8080, by default, and responds to requests at the root URL ("/").
// When a request is received, it prints a greeting message to the console and sends a response back to the client.
//
// Usage:
// Run the program and navigate to the link in web browser.

package main

import (
	"log"
	"fmt"
	"net/http"
	"os"
)

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variable %s is not set.", key)
	}
	return value
}

func main() {
	// Get config from .env
	// Optional env var
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Required env var
	secretMessage := getRequiredEnv("SECRET_MESSAGE")

	// Endpoint response for '/'
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello 'Go' World!! This is a HTTP response.\nThe secret message is: %s", secretMessage) 
	})

	// Endpoint response for '/about'
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This page is about me trying to learn Go and Docker at the sametime! Am I addicted to learning multiple things at once ??")
	})

	// Open server port for requests
	fmt.Printf("Listening on :\nhttp://localhost:%s", port)
	http.ListenAndServe("0.0.0.0:" + port, nil)
}