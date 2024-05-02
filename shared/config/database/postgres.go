package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	// Connect to DB
	stringDB := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"))

	db, err := sql.Open("postgres", stringDB)

	if err != nil {
		return nil, err
	}

	errConn := db.Ping()
	if errConn != nil {
		err = fmt.Errorf("failed to connect to db: %s", err)
		return nil, err
	}

	fmt.Println("successfully connected to db")
	return db, err

}
