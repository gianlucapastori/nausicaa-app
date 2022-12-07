package http

import (
	"github.com/gianlucapastori/nausicaa-app/internal/modules/user"
	"github.com/gorilla/mux"
)

func Init(mux *mux.Router, controller user.Controller) {
	mux.HandleFunc("/auth/google/login", controller.OAuthGoogleLogin()).Methods("GET")
	mux.HandleFunc("/auth/google/callback", controller.OAuthGoogleCallback()).Methods("GET")
}
