package main

import (
	"database/sql"
	"log"

	"tutorial.sqlc.dev/app/api"
	db "tutorial.sqlc.dev/app/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(":8080")
	if err != nil {
		log.Fatal("cannot start server: " + err.Error())
	}
}
