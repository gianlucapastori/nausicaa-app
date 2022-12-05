package auth

type Services interface {
	OAuthGoogleLogin() error
	OAuthGoogleCallback() error
}
