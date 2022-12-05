package user

import "net/http"

type Controller interface {
	OAuthGoogleLogin() http.HandlerFunc
	OAuthGoogleCallback() http.HandlerFunc
}
