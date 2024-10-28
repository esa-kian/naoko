package services

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// DB holds the active database connection
var DB *sql.DB

// ConnectDatabase connects to a specified database
func ConnectDatabase(dbType, host, port, user, password, dbName string) error {
	var dsn string

	switch dbType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	default:
		return fmt.Errorf("unsupported database type: %s", dbType)
	}

	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return err
	}

	// Ping to confirm the connection is active
	if err := db.Ping(); err != nil {
		return err
	}

	// Assign the active connection to the global variable
	DB = db
	return nil
}
