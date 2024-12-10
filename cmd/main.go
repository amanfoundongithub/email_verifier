package main

import (
	"github.com/amanfoundongithub/email_verifier/api"
)

func main() {
	
	// Create new server
	server := api.CreateEmailServer()

	// Activate the server 
	api.ActivateServer(server, ":8080")

}