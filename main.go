package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"simple_bank/api"
	db "simple_bank/db/sqlc"
	"simple_bank/utils"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
