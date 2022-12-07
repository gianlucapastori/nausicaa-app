package services

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/gianlucapastori/nausicaa-app/internal/modules/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authService struct{}

func New() auth.Services {
	return &authService{}
}

var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://172.22.130.38:2323/api/v1/user/auth/google/callback",
	ClientID:     "",
	ClientSecret: "",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const (
	oAuthGoogleUrlApi = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

func (s *authService) OAuthGoogleLogin(w http.ResponseWriter) (string, error) {
	state := s.GenerateStateCookie(w)
	u := googleOAuthConfig.AuthCodeURL(state)

	return u, nil
}

func (s *authService) OAuthGoogleCallback() error {
	log.Panic("unim")
	return nil
}

func (s *authService) GenerateStateCookie(w http.ResponseWriter) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: time.Now().Add(20 * time.Minute)}
	http.SetCookie(w, &cookie)

	return state
}
