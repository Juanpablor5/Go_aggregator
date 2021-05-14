package lealdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var instance *sql.DB

func GetDB() *sql.DB {
	if instance == nil {
		log.Println("creating instance of repository")
		instance = initDb()
	}
	return instance
}

func initDb() *sql.DB {
	log.Println("connecting to db")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", connStr)

	connPoolSize, err := strconv.Atoi(os.Getenv("CONN_POOL_SIZE"))
	if err != nil {
		panic(fmt.Errorf("could not crate the connection pool, %v", err))
	}

	db.SetMaxIdleConns(connPoolSize)
	db.SetMaxOpenConns(connPoolSize)
	db.SetConnMaxLifetime(120 * time.Minute)
	log.Println("connected to db")
	return db
}

// Close closes the database connection
func Close() {
	instance.Close()
}
