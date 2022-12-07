package http

import (
	"fmt"
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
		u, err := c.services.OAuthGoogleLogin(w)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "could not redirect to oauth google")
			return
		}

		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
		return
	}
}

func (c *userController) OAuthGoogleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.services.OAuthGoogleCallback()
		return
	}
}
