package main

import (
	"fmt"
	"log"
	"net/http"
)

// homeHandler handles requests to the root URL ("/").
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request.
	log.Printf("Received request for: %s from %s", r.URL.Path, r.RemoteAddr)

	// Set the Content-Type header to indicate HTML.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write a simple HTML response to the client.
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Go Server</title>
			<style>
				body {
					font-family: 'Inter', sans-serif;
					display: flex;
					justify-content: center;
					align-items: center;
					min-height: 100vh;
					margin: 0;
					background-color: #f0f4f8;
					color: #333;
				}
				.container {
					background-color: #ffffff;
					padding: 40px;
					border-radius: 12px;
					box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
					text-align: center;
					max-width: 500px;
					width: 90%%;
				}
				h1 {
					color: #2c3e50;
					margin-bottom: 20px;
					font-size: 2.5em;
				}
				p {
					font-size: 1.1em;
					line-height: 1.6;
					color: #555;
				}
				.footer {
					margin-top: 30px;
					font-size: 0.9em;
					color: #777;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Welcome to Your Go Server!</h1>
				<p>This is a basic HTTP server written in Go. It's ready to serve your requests.</p>
				<p>You can expand this to build web applications, APIs, or anything you need!</p>
				<div class="footer">
					<p>Server running on port 8080</p>
				</div>
			</div>
		</body>
		</html>
	`)
}

// main function is the entry point of the Go program.
func main() {
	// Register the homeHandler to handle requests for the root URL ("/").
	http.HandleFunc("/", homeHandler)

	// Define the port to listen on.
	port := ":8080"
	log.Printf("Starting server on port %s", port)

	// Start the HTTP server. ListenAndServe blocks until the server stops or an error occurs.
	// If it returns an error, log it.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
