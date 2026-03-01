package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
if we use global variable DB then we have to close it manually by calling db.Close() in main.go
db.Close() is definded at the bottom of the this file (commented)
*/

// var DB *sql.DB

/*
we can call LoadConfig() in main.go and pass it to the ConnectDB() function
Also we can call inside the ConnectDB() function
*/

var config = LoadConfig()

func ConnectDB() *sql.DB {
	// config := LoadConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error While connecting DB", err)
	}
	// Verify connection with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// tune connection pool
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error while connecting with DB", err)
	}

	fmt.Println("DB connected successfully")
	return db
}

// func CloseDB() {
// 	if DB != nil {
// 		if err := DB.Close(); err != nil {
// 			log.Fatalf("Error closing DB connection: %v\n", err)
// 		} else {
// 			fmt.Println("DB connection closed successfully")
// 		}
// 	}
// }
