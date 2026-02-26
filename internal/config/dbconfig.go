package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// var DB *sql.DB
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

	// var err error
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error While connecting DB", err)
	}
	// Verify connection with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

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
