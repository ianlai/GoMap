package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	DBUser     = "gomap_admin"
	DBPassword = "gomap_admin"
	DBName     = "gomap_db"
	DBHost     = "db"
)

type DB struct {
	*sql.DB
}

func NewDBRepo(db *sql.DB) *DB {
	return &DB{
		db,
	}
}

func InitDB() *DB {
	connStr := fmt.Sprintf("sslmode=disable host=%s user=%s dbname=%s password=%s",
		DBHost, DBUser, DBName, DBPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("DB connection successful!!")
	return NewDBRepo(db)
}
