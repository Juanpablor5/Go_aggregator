package facturasdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

var instance *sql.DB

// GetDB returns the database single instance
func GetDB() *sql.DB {
	if instance == nil {
		log.Println("creating instance of facturasdb")
		instance = initDb()
	}
	return instance
}

func initDb() *sql.DB {
	log.Println("connecting to db")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("FACTURAS_DB_USER"),
		os.Getenv("FACTURAS_DB_PWD"),
		os.Getenv("FACTURAS_DB_HOST"),
		os.Getenv("FACTURAS_DB_PORT"),
		os.Getenv("FACTURAS_DB_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("could not create connection to DB, %v", err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	connPoolSize, err := strconv.Atoi(os.Getenv("CONN_POOL_SIZE"))
	if err != nil {
		panic(fmt.Errorf("could not create the connection pool, %v", err))
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
