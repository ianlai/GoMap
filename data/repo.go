package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", "user=gomap_user password=gomap_user dbname=gomap_db sslmode=disable")
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
