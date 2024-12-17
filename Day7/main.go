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
	"os"
	"time"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Record endpoint access in database
		timestamp := time.Now().Format(time.RFC3339) + "\n"
		os.MkdirAll("/data", 0755)
		f, _ := os.OpenFile("/data/visits.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.WriteString(timestamp)
		f.Close()

		// Read and display access logs
		data, err := os.ReadFile("/data/visits.txt")
		if err != nil {
			fmt.Fprintf(w, "Error reading access logs!!")
			return
		}

		fmt.Fprintf(w, "New visit recorded at %s\n\nAll visits :\n%s", timestamp, string(data))
	})

	// response when endpoint '/about' is requested
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This page is about me trying to learn Go and Docker at the sametime! Am I addicted to learning multiple things at once ??")
	})

	// Server port
	fmt.Println("Listening on :\nhttp://localhost:8080")
	http.ListenAndServe(":8080", nil)
}