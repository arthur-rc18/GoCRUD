package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectionPostgres() *sql.DB {

	dbConnection := "user=postgres dbname=crudDB password=zeu$@2022 sslmode=disable"

	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		log.Panic(err.Error())
	}

	return db
}
