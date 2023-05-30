package security

import (

	//"errors"
	"net/http"
	"time"

	//"net/url"

	"blow.com/bogoso.backend/app"
	//"../utils"
	//"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	//"strings"
	"github.com/dgrijalva/jwt-go/request"
	log "github.com/sirupsen/logrus"
)

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(email string, duration int) (string, error) {
	logger := log.WithFields(log.Fields{"module": "security.CreateToken"})
	privateKey := app.PrivateKey
	//    fmt.Println("private key: ", string(mySigningKey)  )

	// Create the Claims

	ut := time.Now().Add(time.Duration(duration) * time.Hour).Unix()
	claims := MyCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: ut,
			Issuer:    "bogoso",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	mySigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		logger.Errorf("signed token error: %S", err)
		return "", err
	}
	return ss, nil
}

func AuthTokenHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.WithFields(log.Fields{"module": "security.AuthTokenHandler"})
		key, err := jwt.ParseRSAPublicKeyFromPEM(app.PublicKey)
		//		fmt.Println("r *http.Request", r)
		token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			//return app.PublicKey, nil
			return key, nil
		})
		if err != nil {
			//app.Logger(err.Error())
			logger.Errorf("ParseFromRequestWithClaims: %s", err)
			http.Error(w, "invalid_token", http.StatusUnauthorized)
			return
		}
		if token.Valid {
			custClaims := token.Claims.(*MyCustomClaims)
			email := custClaims.Email
			logger.Infof("token.Email: %s", email)
			r.Header.Set("email", email)
			next.ServeHTTP(w, r)
		} else {
			//	app.Logger("Invalid Token Received")
			http.Error(w, "Invalid Session ", http.StatusUnauthorized)
		}
	})
}

func ValidateToken(tokenString string) (string, error) {
	logger := log.WithFields(log.Fields{"module": "security.ValidateToken"})
	logger.Infof("tokenString :%s", tokenString)
	publicKey, err := jwt.ParseRSAPrivateKeyFromPEM(app.PublicKey)
	if err != nil {
		logger.Errorf("publicKey :", err)
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		logger.Infof("%v %v", claims.Email, claims.StandardClaims.ExpiresAt)
		return claims.Email, nil
	} else {
		logger.Errorf("Error validating claims: %s", err)
		return "", err
	}
}
