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

	var dbUrl = "libsql://testing-turso-bubemi054.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDI0LTAxLTE0VDIwOjA1OjA1LjIwODU4MDY2MloiLCJpZCI6IjNlNDA4N2YxLWExYmUtMTFlZS04ZTNiLTgyZDViNGZlYjEyYyJ9.Mduq3egH_v29VOZzaGOTWyeMOfLAG538Z_pLx6vwRtupOSwYYZ2J2unXiz23R4XYLKyxqQO3tk6YN8I_auUYBA"
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
	fmt.Println("table creation successful")
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
