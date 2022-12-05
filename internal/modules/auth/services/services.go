package services

import (
	"log"

	"github.com/gianlucapastori/nausicaa-app/internal/modules/auth"
)

type authService struct{}

func New() auth.Services {
	return &authService{}
}

func (s *authService) OAuthGoogleLogin() error {
	log.Panic("unim")
	return nil
}

func (s *authService) OAuthGoogleCallback() error {
	log.Panic("unim")
	return nil
}
