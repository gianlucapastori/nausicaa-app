package auth

import "net/http"

type Services interface {
	OAuthGoogleLogin(http.ResponseWriter) (string, error)
	OAuthGoogleCallback() error
	GenerateStateCookie(http.ResponseWriter) string
}
