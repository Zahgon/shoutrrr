package smtp

import (
	"net/smtp"
)

type oauth2Auth struct {
	username, accessToken string
}

// OAuth2Auth returns an Auth that implements the SASL XOAUTH2 authentication
// as per https://developers.google.com/gmail/imap/xoauth2-protocol
func OAuth2Auth(username, accessToken string) smtp.Auth {
	_ = "STUB: not implemented"
	return *new(smtp.Auth)
}

func (a *oauth2Auth) Start(_ *smtp.ServerInfo) (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (a *oauth2Auth) Next(_ []byte, _ bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
