package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	mux *mux.Router
	db  *sqlx.DB
}

func New(db *sqlx.DB) *Server {
	return &Server{mux: mux.NewRouter(), db: db}
}

var (
	port int = 2323
)

func (s *Server) Serve() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	if err := s.Init(); err != nil {
		return err
	}

	log.Println("-> server listening on port", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	log.Println("\n-> server shutting down...")
	return server.Shutdown(ctx)
}
