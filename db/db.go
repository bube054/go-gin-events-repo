package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func InitDB() error {
	var err error

	dbUrl := os.Getenv("TURSO_DB_URL")

	if dbUrl == "" {
		return errors.New("err loading TURSO_DB_URL from .env")
	}

	dbToken := os.Getenv("TURSO_AUTH_TOKEN")

	if dbToken == "" {
		return errors.New("err loading TURSO_AUTH_TOKEN from .env")
	}

	fmt.Println("TURSO_DB_URL: ", dbUrl)
	fmt.Println("Token: ", dbToken)

	var fmtDBUrl = dbUrl + "?authToken=" + dbToken
	DB, err = sql.Open("libsql", fmtDBUrl)
	if err != nil {
		panic(err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	fmt.Println("Connection to db successful")

	CreateTables()

	return nil
}

func CreateTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not create users table.")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id) 
		)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err.Error())
		panic("Could not create events table.")
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(event_id) REFERENCES events(id) 
		)
	`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		fmt.Println(err.Error())
		panic("Could not create registrations table.")
	}

	fmt.Println("All table(s) creation successful")
}

func DropTables() {
	dropUsersTable := `
		DROP TABLE events
	`
	_, err := DB.Exec(dropUsersTable)
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not drop users table.")
	}
}
