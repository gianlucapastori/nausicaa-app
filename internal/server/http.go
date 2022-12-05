package server

import (
	"fmt"
	"net/http"
)

func (s *Server) Init() error {
	vRoute := s.mux.PathPrefix("/api/v1").Subrouter()

	vRoute.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "server up")
		return
	})

	return nil
}
