package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB 

func InitDB() {
	var err error
	dsn := "root:root@tcp(localhost:3306)/student"

	DB, err = sql.Open("mysql", dsn) 
	if err != nil {
		log.Fatal("Error connecting:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Connected to MySQL successfully!")
}
