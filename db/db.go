package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var Driver *sql.DB
var once sync.Once

func init() {
	// database URL
	// please change the bd url befor you test
	dbUrl := "postgres://postgres@localhost:5432/testdb?sslmode=disable"

	// open the database
	pgDriver, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Errorf("error opening db connection: ", err)
	}

	pgDriver.SetMaxOpenConns(1)
	pgDriver.SetMaxIdleConns(1)
	pgDriver.SetConnMaxLifetime(time.Hour)

	Driver = pgDriver
}
