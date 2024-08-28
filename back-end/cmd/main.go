package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitub.com/matheus-hrm/curiously/cmd/api"
)

func main() {

	db, err := pgxpool.New(context.Background(),
		`host=localhost 
		port=5433
		user=mou
		dbname=curious_db
		password=12345	
		sslmode=disable`,
	)
	if err != nil {
		log.Fatal(err)
	}
	InitStorage(db)
	defer db.Close()

	server := api.New(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func InitStorage(db *pgxpool.Pool) {
	err := db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
}