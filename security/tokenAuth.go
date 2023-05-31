package security

import (
	log "github.com/sirupsen/logrus"
)

// Example User Authentication shows how a typical application can verify a login attempt
func TokenAuth(p_userid string) (UserProfile, error) {
	logger := log.WithFields(log.Fields{"module": "security.BasicAuth"})
	// The username and password we want to check
	username := p_userid
	logger.Infof("username concat: %s", username)
	var u UserProfile

	token, err := CreateToken(p_userid, 24)
	if err != nil {
		logger.Errorf("CreateToken: %s", err)
		return u, err
	}

	u.UserName = username
	u.Token = token
	return u, nil
}
