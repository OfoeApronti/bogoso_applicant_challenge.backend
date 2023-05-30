package hdl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"blow.com/bogoso.backend/app"

	//"blow.com/bogoso.backend/security"

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

	applicant_name := ""
	phone := ""
	file_name := ""
	var id int
	err = app.DBBogoso.QueryRow(`SELECT id, applicant_name, phone,file_name
	FROM bogoso.cv_files where email=$1`, e.Email).Scan(&id, &applicant_name, &phone, &file_name)
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	id_string := fmt.Sprint(id)
	payload := CV_Payload{id_string, applicant_name, phone, file_name}
	js, err := json.Marshal(payload)
	if err != nil {
		logger.Errorf("json.Unmarshal(body, &t): %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
})