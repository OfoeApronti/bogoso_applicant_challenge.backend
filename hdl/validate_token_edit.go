package hdl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"blow.com/bogoso.backend/app"

	"blow.com/bogoso.backend/security"

	//"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type TokenData struct {
	Email string `json:"email" validate:"required"`
	Token string `json: "token"  validate:"required"`
}

type CV_Payload struct {
	Id            string `json:"id"`
	ApplicantName string `json:"applicant_name"`
	Phone         string `json:"phone"`
	FileName      string `json:"file_name"`
	Email         string `json:"email"`
}

var Handler_Validate_Token = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.Handler_Validate_Token"})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("ioutil.ReadAll(r.Body): %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	logger.Info("Handler_Validate_Token")
	var e TokenData
	err = json.Unmarshal(body, &e)
	if err != nil {

		logger.Errorf("json.Unmarshal(body, &t): %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}

	var email string
	fmt.Println(e.Email, e.Token)
	err = app.DBBogoso.QueryRow(`SELECT email
	FROM bogoso.temp_token where email=$1 and token=$2`, e.Email, e.Token).Scan(&email)
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	if email == "" {
		logger.Error("Invalid token provided")
		http.Error(w, "Invalid token provided", http.StatusInternalServerError)
		return
	}

	uPayload, err := security.TokenAuth(e.Email)
	if err != nil {
		logger.Errorf("Error validating token: %s ", e.Email)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(uPayload)
	if err != nil {
		logger.Errorf("Error marshalling: %s", e.Email)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
})
