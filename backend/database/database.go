package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

func Database(dbName string) (*sql.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("ADDR"), // maybe create a map between different ports in the future
		DBName: dbName,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		err = fmt.Errorf("database.Database.error: %s", err)
	}
	return db, err
}
