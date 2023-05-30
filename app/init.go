package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//_ "github.com/godror/godror"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type AccessConfig struct {
	Bogoso_host         string `json:"bogoso_db_host"`
	Bogoso_port         int    `json:"bogoso_db_port"`
	Bogoso_user         string `json:"bogoso_db_user"`
	Bogoso_password     string `json:"bogoso_db_pwd"`
	Bogoso_dbname       string `json:"bogoso_db_name"`
	Bogoso_schema       string `json:"bogoso_db_schema"`
	PortalAdmin         string `json:"portal_admin"`
	PortalAdminEmail    string `json:"portal_admin_email"`
	PortalAdminPassword string `json:"portal_admin_password"`
	GmailAppToken string `json:"gmail_app_token"`
}

var (
	signedKey           string
	PrivateKey          []byte
	PublicKey           []byte
	DBBogoso            *sql.DB
	bogoso_host         string
	bogoso_port         int
	bogoso_user         string
	bogoso_password     string
	bogoso_dbname       string
	bogoso_schema       string
	LogFormat           *logrus.TextFormatter
	PortalAdmin         string
	PortalAdminEmail    string
	PortalAdminPassword string
	GmailAppToken string
)

func Init() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	LogFormat = customFormatter
	log.SetFormatter(LogFormat)
	logger := log.WithFields(log.Fields{"module": "app.Init"})
	logger.Info("Initialized and loading parameters")
	conf := getAccessConfigs()

	PrivateKey, PublicKey = getKey()
	signedKey = "secret"
	bogoso_dbname = conf.Bogoso_dbname
	bogoso_host = conf.Bogoso_host
	bogoso_password = conf.Bogoso_password
	bogoso_port = conf.Bogoso_port
	bogoso_schema = conf.Bogoso_schema
	bogoso_user = conf.Bogoso_user
	PortalAdmin = conf.PortalAdmin
	PortalAdminEmail = conf.PortalAdminEmail
	PortalAdminPassword = conf.PortalAdminPassword
	GmailAppToken =conf.GmailAppToken
	DBBogoso = setupDbBogoso()
}
func getKey() (a []byte, b []byte) {
	privateKey, _ := ioutil.ReadFile("demo.rsa")
	publicKey, _ := ioutil.ReadFile("demo.pub.rsa")
	return privateKey, publicKey
}

func getAccessConfigs() AccessConfig {
	logger := log.WithFields(log.Fields{"module": "app.getAccessConfigs"})
	logger.Info("Fetching environment variables")

	var Bk AccessConfig
	//	var Ret []string
	raw, err := ioutil.ReadFile("./access.config")
	if err != nil {
		logger.Error("Error reading the access.json file")
		logger.Fatal(err.Error())

	}
	err = json.Unmarshal(raw, &Bk)
	if err != nil {
		logger.Error("Error marshalling the access.json file")
		logger.Fatal(err.Error())
	}
	logger.Info("Completed fetching environment variables")
	return Bk
}

func setupDbBogoso() *sql.DB {
	logger := log.WithFields(log.Fields{"module": "app.setupDbBogoso"})
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		bogoso_host, bogoso_port, bogoso_user, bogoso_password, bogoso_dbname, bogoso_schema)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Fatalf("setupDbBogoso: %s", err) // nil means it is a match
	}
	return db
}
