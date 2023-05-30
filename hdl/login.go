package hdl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"blow.com/bogoso.backend/security"

	//"fmt"
	log "github.com/sirupsen/logrus"
)

type loginValues struct {
	Userid   string `json:"user_id"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.Casa"})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error reading login request object: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(body))
	var t loginValues
	err = json.Unmarshal(body, &t)
	if err != nil {
		logger.Errorf("Error reading login request object: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uPayload, err := security.BasicAuth(t.Userid, t.Password)
	if err != nil {
		logger.Errorf("Error validating user: %s ", t.Userid)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(uPayload)
	if err != nil {
		logger.Errorf("Error resolving LDAP payload for: %s", t.Userid)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	//http.Redirect(w,r,"http://10.179.100.152:3001/static/index.html",http.StatusMovedPermanently)
}
