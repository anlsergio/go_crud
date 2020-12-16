package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL implicit connection driver
)

// Connect open up the database connection
func Connect() (*sql.DB, error) {
	connectionString := "devbook:AReallyStrongPassword@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err // where "nil" is the "empty" or 0 value of the expected return of "*sql.DB"
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}