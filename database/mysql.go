package database

import (
	"database/sql"
	"log"
)

var status string

func ConnectMysql() string {

	// Replace with your database credentials
	db, err := sql.Open("mysql", "root:1T$hutt!ers@tcp(localhost:3306)/itsm")
	if err != nil {
		log.Fatal(err)
		status = "Unable to connect to mysql database"
	}
	status = "Connected to the mysql database"
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return status
}
