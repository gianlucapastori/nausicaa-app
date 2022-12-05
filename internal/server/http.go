package server

import (
	"fmt"
	"net/http"

	auth_service "github.com/gianlucapastori/nausicaa-app/internal/modules/auth/services"
	user_http "github.com/gianlucapastori/nausicaa-app/internal/modules/user/http"
)

func (s *Server) Init() error {
	// repo dependency injection

	// service dependency injection
	authService := auth_service.New()

	// controller dependecy injection
	userController := user_http.New(authService)

	vRoute := s.mux.PathPrefix("/api/v1").Subrouter()
	userRoute := vRoute.PathPrefix("/user").Subrouter()

	vRoute.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "server up")
		return
	})

	user_http.Init(userRoute, userController)

	return nil
}
