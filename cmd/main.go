package main

import (
	"net/http"

	"github.com/amanfoundongithub/email_verifier/api"
)

func main() {
	
	// Create new server
	server := api.CreateEmailServer()

	//
	http.ListenAndServe(":8080", server) 

}