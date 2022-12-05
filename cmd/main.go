package main

import (
	"log"

	"github.com/gianlucapastori/nausicaa-app/internal/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "__nausica.db")
	if err != nil {
		log.Panicln(err)
	}

	s := server.New(db)

	if err := s.Serve(); err != nil {
		log.Panicln(err)
	}
}
