package main

import (
	"fmt"
	"net/http" // Provides the tools needed to build an HTTP server and handle requests.
)

func main() {
	// Define a handler function for the root endpoint ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // A handler is a function that gets called when someone accesses this URL.
		// The anonymous function (func(w http.ResponseWriter, r *http.Request)) does two things:
		// w (ResponseWriter): Used to send the response back to the client.
		// r (Request): Represents the incoming HTTP request.
		fmt.Fprintln(w, "Hello, World!") // Write response to the client
		// Sends the text "Hello, World!" back as a response to the client.

	})

	// Define a handler function for another endpoint ("/greet")
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) { // http://localhost:8080/greet?name=Alice
		name := r.URL.Query().Get("name") // Get query parameter "name"
		// Reads the query parameter name from the URL (e.g., ?name=Alice means name is Alice)
		if name == "" {
			name = "Guest" // Default name if not provided
		}
		fmt.Fprintf(w, "Hello, %s!", name) // Write personalized greeting
		// Sends a personalized greeting back to the client. If the name is "Alice," the response will be "Hello, Alice!".
	})

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // Starts an HTTP server on port 8080.
	// The second argument nil means the server uses the default handler.
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// ctrl + shift + c to copy the http adress from the terminal.
// ctrl + c to stop the running programm.

/* How the Handlers Work:

1) http.HandleFunc("/"): ex: http://localhost:8080

This defines a handler for the root path (/).
It's invoked if the request URL is / or if there are no more specific routes that match the request.

2) http.HandleFunc("/greet"): ex: http://localhost:8080/greet (in terminal only "http://localhost:8080" this is given), we have to put "/greet" by ourself.
// or if we want to display name instead of guest(default name), then we can "http://localhost:8080/greet?name=iman" search this web page also.

This defines a handler for the /greet path.
If someone sends a request to /greet or /greet?name=Alice, this handler is invoked instead of the root (/) handler. */
