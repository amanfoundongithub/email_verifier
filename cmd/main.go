package main

import (
	"fmt"
	"github.com/amanfoundongithub/email_verifier/verifier"
)

func main() {
	email := "billygates124356@gmail.com"

	if verifier.IsValidRegex(email) {
		fmt.Println("syntax checked, now validity...")
		
		response, err := verifier.VerifyDomain(email)

		if err != nil {
			fmt.Println(err)
		} else {
			if response {
				fmt.Println("Verified")
			} else {
				fmt.Println("Unverified")
			}
		}
	} else {
		fmt.Println("yes")
	}
}