package hdl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"blow.com/bogoso.backend/app"

	//"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// type TokenData struct {
// 	Email string `json:"email" validate:"required"`
// 	Token string `json: "token"  validate:"required"`
// }

// type CV_Payload struct {
// 	Id            string `json:"id"`
// 	ApplicantName string `json:"applicant_name"`
// 	Phone         string `json:"phone"`
// 	FileName      string `json:"file_name"`
// 	Email         string `json:"email"`
// }

type IdData struct {
	ID string `json:"id"`
}

var Handler_GetSingle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.Handler_GetSingle"})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("ioutil.ReadAll(r.Body): %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	var e IdData
	err = json.Unmarshal(body, &e)
	if err != nil {

		logger.Errorf("json.Unmarshal(body, &t): %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}

	var email, applicant_name, phone, file_name string
	fmt.Println(e.ID)
	err = app.DBBogoso.QueryRow(`SELECT applicant_name, email, phone, file_name
	FROM bogoso.cv_files where id=$1`, e.ID).Scan(&applicant_name, &email, &phone, &file_name)
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	uPayload := CV_Payload{e.ID, applicant_name, phone, file_name, email}
	js, err := json.Marshal(uPayload)
	if err != nil {
		logger.Errorf("Error marshalling: %s", e.ID)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
})
