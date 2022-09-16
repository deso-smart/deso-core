package main

import (
	"github.com/deso-smart/deso-core/v3/lib"
	"github.com/go-pg/pg/v10"
	"log"
	"os"

	"github.com/deso-smart/deso-core/v3/migrate"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

const directory = "migrate"

//
// This file provides a database migration CLI. For example:
//
// Create a new migration: go run migrate.go create MyMigration
// Migrate the database:   go run migrate.go migrate
// Rollback the database:  go run migrate.go rollback
//

func main() {
	migrate.LoadMigrations()

	// Default postgres connection options
	pgOptions := &pg.Options{
		Addr:     "localhost:5432",
		User:     "admin",
		Database: "admin",
		Password: "",
	}

	// If set, use custom postgres connection options.
	if len(os.Getenv("POSTGRES_URI")) > 0 {
		pgOptions = lib.ParsePostgresURI(os.Getenv("POSTGRES_URI"))
	}

	db := pg.Connect(pgOptions)
	err := migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
