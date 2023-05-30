package security

import (

	//"errors"
	//"net/http"
	"blow.com/bogoso.backend/app"
	//"../utils"
	//"encoding/json"
	jwt "gopkg.in/dgrijalva/jwt-go.v2"
	//"strings"
	//"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func CreateTokenOld() (string, error) {
	logger := log.WithFields(log.Fields{"module": "security.token"})
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	//fmt.Println(app.PrivateKey)
	tokenString, err := token.SignedString(app.PrivateKey)
	if err != nil {
		logger.Errorf("Err CreateToken: %s", err)
		return "", err
	}
	return tokenString, nil
}
