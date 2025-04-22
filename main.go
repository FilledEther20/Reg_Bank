package main

import (
	"database/sql"
	"log"

	"github.com/FilledEther20/Reg_Bank/api"
	"github.com/FilledEther20/Reg_Bank/db/sqlc"
	"github.com/FilledEther20/Reg_Bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := sqlc.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Problem to start with server ", err)
	}
}
