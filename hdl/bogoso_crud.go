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

type EditData struct {
	Id    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

var BogosoDeleteCV = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.BogosoDeleteCV"})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("ioutil.ReadAll(r.Body): %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	logger.Info("BogosoDeleteCV")
	var e EditData
	err = json.Unmarshal(body, &e)
	if err != nil {

		logger.Errorf("json.Unmarshal(body, &t): %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}

	_, err = app.DBBogoso.Exec(`delete from bogoso.cv_files
			where id=$1`, e.Id)
	if err != nil {

		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	fmt.Fprintf(w, "Delete successful")
})

var ValidateToken = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.GetElevyRecords"})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("ioutil.ReadAll(r.Body): %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}
	var e EditData
	err = json.Unmarshal(body, &e)
	if err != nil {

		logger.Errorf("json.Unmarshal(body, &t): %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//		panic(err)
	}

	_, err = app.DBBogoso.Exec(`delete from bogoso.cv_files
			where id=$1`, e.Id)
	fmt.Fprintf(w, "Delete successful")
})
