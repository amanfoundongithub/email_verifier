package api

import (
	"encoding/json"
	"net/http"
	"github.com/amanfoundongithub/email_verifier/verifier"
)

// Request Body
type RequestBody struct {
	Email string `json:"email"`
}

// Response Body 
type ResponseBody struct {
	Message string  `json:"message"`
	Success bool  `json:"is_valid"`
	Email   string  `json:"email"`
}



func emailHandler(w http.ResponseWriter, r *http.Request) {
	
	// Check method if it is POST or NOT
	if r.Method != http.MethodPost {
		http.Error(w, "ERR_METHOD_NOT_ALLOWED_ONLY_POST", http.StatusMethodNotAllowed)
		return 
	}

	// Body parsing
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "ERR_INVALID_JSON_BODY", http.StatusUnprocessableEntity)
		return  
	}

	// Defer till the end
	defer r.Body.Close()

	// Get email
	email := requestBody.Email

	message := ""
	is_valid := true

	// Verified email?
	verified_regex := verifier.IsValidRegex(email)

	// Update Message 
	if !verified_regex {
		message = InvalidRegexMessage
		is_valid = false 
	} else {
		verified_domain, err := verifier.VerifyDomain(email)

		// 500
		if err != nil {
			http.Error(w, "ERR_INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
			return 
		}

		if !verified_domain {
			message = InvalidDomainMessage
			is_valid = false 
		} else {
			message = ValidEmailMessage
			is_valid = true 
		}
	}

	// Define the output response
	response := ResponseBody{
		Message: message,
		Success: is_valid,
		Email: email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) 


}

func addEmailHandlerToMUX(mux * http.ServeMux, route string) {
	
	// Add the email Handler to the mux
	mux.HandleFunc(route, emailHandler)

}

