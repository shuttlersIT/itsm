package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var status string

func ConnectMysql() (string, *sql.DB) {

	// Replace with your database credentials
	db, err := sql.Open("mysql", "root:1T$hutt!ers@tcp(localhost:3306)/itsm")
	if err != nil {
		log.Fatal(err)
		status = "Unable to connect to mysql database"
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Check if the "tickets" table exists
	if TableExists(db, "tickets") {
		fmt.Println("The 'tickets' table exists.")
	} else {
		fmt.Println("The 'tickets' table does not exist.")
	}

	return status, db
}

// Create a table to store tickets in the database
func CreateTicketsTable(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS tickets (
            id INT AUTO_INCREMENT PRIMARY KEY,
			first_name    VARCHAR(100) NOT NULL,
			last_name     VARCHAR(100) NOT NULL,
			staff_email   VARCHAR(100) NOT NULL,
			username     VARCHAR(50) NOT NULL,
			position_iD   INT,
			department_iD INT,
    		FOREIGN KEY (position_id) REFERENCES positions(id),
    		FOREIGN KEY (department_id) REFERENCES departments(id)
        )`
	_, err := db.Exec(query)
	return err
}

// Check if a table exists in the database
func TableExists(db *sql.DB, tableName string) bool {
	query := "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = ?"
	var count int
	err := db.QueryRow(query, tableName).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}
