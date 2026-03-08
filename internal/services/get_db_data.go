package services

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID       int
	UserName string
	Email    string
}

func GetDBData(db *sql.DB, lastId int) ([]User, error) {
	rows, err := db.Query(`
	SELECT ID, userName, email 
	FROM users 
	WHERE ID > ? 
	ORDER BY ID ASC
	LIMIT 2`, lastId)
	if err != nil {
		fmt.Printf("[%s] Error: %v\n", time.Now().Format(time.RFC3339), err)
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
			// fmt.Println("Error while scanning", err)
			return nil, err
		}
		// fmt.Println(u)
		users = append(users, u)
	}
	// fmt.Println("rows", rows)
	return users, nil
}
