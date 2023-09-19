package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials

	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		// send error back
		app.errorLog.Println("invalid json")
		payload.Error = true
		payload.Message = "Invalid JSON"

		out, err := json.MarshalIndent(payload, "", "\t")

		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}

	// authenticate
	app.infoLog.Println(creds.UserName, creds.Password)

	// send back a response
	payload.Error = false
	payload.Message = "Login Successful"

	out, err := json.MarshalIndent(payload, "", "\t")

	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
