package main

import (
	"database/sql"
	"log"

	"github.com/FilledEther20/Reg_Bank/api"
	"github.com/FilledEther20/Reg_Bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := sqlc.NewStore(conn)
	server := api.NewServer(&store)
	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Problem to start with server ", err)
	}
}
