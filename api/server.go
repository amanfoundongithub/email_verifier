package api

import (
	"fmt"
	"net/http"
)

type EmailServer http.Handler

func CreateEmailServer() EmailServer {

	// Initialize a router 
	mux := http.NewServeMux()

	// Add handler to the mux
	addEmailHandlerToMUX(mux,"/verify/email") 

	// Add handler to it
	middleware_handler := addMiddlewareToMUX(mux) 

	return middleware_handler

}

func ActivateServer(server EmailServer, port string) {

	// Activate server
	http.ListenAndServe(port, server) 

	// Print the statement 
	fmt.Printf("Server is active on port: %v", port) 
}
