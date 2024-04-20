package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func InitDB() {
	var err error

	var dbUrl = "libsql://events-scheduler-bube054.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MTM1ODQ3OTAsImlkIjoiMDYzNDQ5MTAtYWNkMS00NTNmLTk2ZjQtNWZkNzc2MjM1M2E5In0.w1U_j064UUzBOI1ZI5v2BlNwqQE_dQR-NFgB90gBOjFVUuqD_WQIygDb2ffzx5Mt3WQZ0r0knlG3Ckp-3t3QBQ"
	DB, err = sql.Open("libsql", dbUrl)
	if err != nil {
		panic(err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	fmt.Println("Connection to db successful")
	// DropTables()
	CreateTables()
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
