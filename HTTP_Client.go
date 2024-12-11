package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// HTTP Server
func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Send a response back to the client
		fmt.Fprintln(w, "Hello from the server!")
	})

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// HTTP Client
func startClient() {
	// Give the server some time to start
	time.Sleep(2 * time.Second)

	// Send a GET request to the server
	resp, err := http.Get("http://localhost:8080")
	// This sends an HTTP GET request to the URL http://localhost:8080.
	//The response will be stored in the resp variable, and any errors during the request are stored in err.
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close() //  This ensures that the response body is closed once you're done reading it.

	// Read the response body
	body, err := io.ReadAll(resp.Body) // Reads the entire response body and stores it in the body variable.
	// The resp.Body is an io.Reader, so you can use io.ReadAll to read all the data at once.
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println("Client received response:", string(body))
}

func main() {
	// Start the HTTP server in a goroutine
	go startServer()

	// Start the HTTP client after a short delay to allow the server to start
	startClient()
	//ser()
}

// here in the HTTP directory, when you want to execute what programm you need to change the main func accordingly,
// because both the HTTP Function cannot work together in one main function.
// whenever you execute what, give main() to that programm.
