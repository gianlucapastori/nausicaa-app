package http

import (
	"net/http"

	"github.com/gianlucapastori/nausicaa-app/internal/modules/auth"
	"github.com/gianlucapastori/nausicaa-app/internal/modules/user"
)

type userController struct {
	services auth.Services
}

func New(services auth.Services) user.Controller {
	return &userController{services: services}
}

func (c *userController) OAuthGoogleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.services.OAuthGoogleLogin()
		return
	}
}

func (c *userController) OAuthGoogleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
