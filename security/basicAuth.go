package security

import (
	"errors"
	"strings"

	"blow.com/bogoso.backend/app"
	log "github.com/sirupsen/logrus"
)

type UserProfile struct {
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

// Example User Authentication shows how a typical application can verify a login attempt
func BasicAuth(p_userid string, p_password string) (UserProfile, error) {
	logger := log.WithFields(log.Fields{"module": "security.BasicAuth"})
	// The username and password we want to check
	username := p_userid
	password := p_password
	logger.Infof("username concat: %s", username)
	var u UserProfile

	if strings.ToLower(p_userid) != strings.ToLower(app.PortalAdminEmail) && strings.ToLower(p_userid) != strings.ToLower(app.PortalAdmin) {
		return u, errors.New("Invalid User")
	}
	if password != app.PortalAdminPassword {
		return u, errors.New("Invalid Password")
	}

	token, err := CreateToken(p_userid, 24)
	if err != nil {
		logger.Errorf("CreateToken: %s", err)
		return u, err
	}

	u.UserName = username
	u.Token = token
	return u, nil
}
