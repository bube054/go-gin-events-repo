package models

import (
	// "fmt"
	// "time"

	"errors"

	"example.com/learning/db"
	"example.com/learning/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return nil
	}

	id, err := result.LastInsertId()

	u.ID = id

	return err
}

func (u User) ValidateCredentials () error {
	query := `SELECT password from users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return  errors.New("Credentials are invalid!")
	}

	passwordIsValid := utils.AreHashedPasswordsEqual(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials are invalid!")
	}

	return nil
}

// func GetAllUsers() ([]User, error) {
// 	query := `SELECT * FROM Users`

// 	rows, err := db.DB.Query(query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var Users []User
// 	fmt.Println("have my Users")

// 	for rows.Next() {
// 		var User User
// 		err := rows.Scan(&User.ID, &User.Name, &User.Description, &User.Location, &User.DateTime, &User.UserID)
// 		if err != nil {
// 			fmt.Println("Error", err.Error())
// 			return nil, err
// 		}
// 		Users = append(Users, User)
// 	}
// 	fmt.Println("Users", Users)
// 	defer rows.Close()

// 	return Users, nil
// }

// func GetUserById (id int64) (User, error) {
// 	query := `SELECT * FROM Users where id = ?`
// 	row := db.DB.QueryRow(query, id)

// 	var User User

// 	err := row.Scan(&User.ID, &User.Name, &User.Description, &User.Location, &User.DateTime, &User.UserID)

// 	if err != nil {
// 		fmt.Println("Err", err.Error())
// 		return User{}, nil
// 	}

// 	return User, nil
// }

// func (e User) Update() error {
// 	query := `
// 		UPDATE Users
// 		SET name = ?, description = ?, location = ?, dateTime = ?
// 		WHERE id = ?
// 	`

// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
// 	return err
// }

// func (e User) Delete() error {
// 	query := `
// 		DELETE FROM Users
// 		WHERE id = ?
// 	`

// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.ID)

// 	return err
// }
