package services

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int
	UserName string
	Email    string
}

func GetDBData(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT ID, userName, email FROM users Limit 10")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		// var ID int
		// var userName string
		// var email string
		if err := rows.Scan(&u.ID, &u.UserName, &u.Email); err != nil {
			fmt.Println("Error while scanning", err)
			return nil, err
		}
		fmt.Println(u)
		users = append(users, u)
	}
	// fmt.Println("rows", rows)
	return users, nil
}
