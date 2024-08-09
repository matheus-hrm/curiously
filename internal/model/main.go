package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5433 user=mou dbname=curious_db password=12345 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	m, err := migrate.NewWithDatabaseInstance(
		"file:///~/code/crlsy/internal/migrations",
		"pgx",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}