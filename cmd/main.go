package main

import (
	"fmt"
	"github.com/amanfoundongithub/email_verifier/verifier/regex"
)

func main() {
	email := "aman@js"

	if regex.isValidRegex(email) {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}
}