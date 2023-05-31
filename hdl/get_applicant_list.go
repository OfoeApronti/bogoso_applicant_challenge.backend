package hdl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"blow.com/bogoso.backend/app"
	//"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

var GetApplicantList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var rw json.RawMessage

	err := app.DBBogoso.QueryRow(`
  select json_agg(t) from ( 
		SELECT id, applicant_name, email, phone, file_name, to_char(created,'yyyy-mm-dd') created_date
		FROM bogoso.cv_files
		) t`).Scan((*[]byte)(&rw))
	if err != nil {
		log.Error(err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rw)
})

var SearchApplicantById = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.SearchApplicantList"})
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
	var rw json.RawMessage

	err = app.DBBogoso.QueryRow(`
  select json_agg(t) from ( 
		SELECT id, applicant_name, email, phone, file_name, to_char(created,'yyyy-mm-dd') created_date
		FROM bogoso.cv_files  where id=$1 order by id asc
		) t`, e.ID).Scan((*[]byte)(&rw))
	if err != nil {
		log.Error(err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rw)
})
