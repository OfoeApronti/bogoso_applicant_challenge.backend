package hdl

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"

	"blow.com/bogoso.backend/app"

	//"blow.com/bogoso.backend/security"
	"blow.com/bogoso.backend/utils"

	//"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type Data struct {
	Email string `json:"email" validate:"required"`
}

var Handler_SendEmail = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.GetElevyRecords"})
	logger.Info("handler send email")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("ioutil.ReadAll(r.Body): %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	var e Data
	err = json.Unmarshal(body, &e)
	if err != nil {

		logger.Errorf("json.Unmarshal(body, &t): %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}

	//first delete the last record
	_, err = app.DBBogoso.Exec(`DELETE FROM bogoso.temp_token
	WHERE email=$1
	`, e.Email)
	if err != nil {

		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}

	token := utils.GetRandomNumber()
	//insert the new data
	_, err = app.DBBogoso.Exec(`INSERT INTO bogoso.temp_token
	(email, "token", created)
	VALUES($1, $2, now())
	`, e.Email, token)
	if err != nil {

		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}

	err = utils.SendEmail(e.Email, token)
	if err != nil {

		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
})
