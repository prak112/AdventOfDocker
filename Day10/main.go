// Package main implements a simple HTTP server that responds with a greeting message.
// 
// The server listens on port 8080 and responds to requests at the root URL ("/").
// When a request is received, it prints a greeting message to the console and sends
// a response back to the client.
//
// Usage:
//   Run the program and navigate to http://localhost:8080 in your web browser to see the response.

package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello 'Go' World!! This is a HTTP response.") 
		// / this will be printed as response when endpoint '/' is requested
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This page is about me trying to learn Go and Docker at the sametime! Am I addicted to learning multiple things at once ??")
		// / this will be printed as response when endpoint '/about' is requested
	})
	fmt.Println("Listening on :\nhttp://localhost:8080")
	http.ListenAndServe(":8080", nil)
}