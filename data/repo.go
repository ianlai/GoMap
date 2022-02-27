package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	DBUser     = "gomap_admin"
	DBPassword = "gomap_admin"
	DBName     = "gomap_db"

	//Cnnect by docker-compose
	//DBHost     = "db"  

	//Connect from host machine through Host IP (confirmed)
	//DBHost = "127.0.0.1" //connect by host
	//DBHost = "localhost" //connect by host
	//DBPort = "5432"

	//Connect from Docker through Internal IP (confirmed)
	//DBHost = "172.27.0.2"
	//DBPort = "5432"
)

// Repo interface
type Repo interface {
	InsertRecord(string, string) error
	GetRecordsSortedByVal(int) ([]Record, error)
}

type DB struct {
	*sql.DB // struct embedding
}

func NewDBRepo(db *sql.DB) *DB {
	return &DB{
		db,
	}
}

func InitDB() *DB {

	connStr := fmt.Sprintf(
		"sslmode=disable host=%s port=%s user=%s dbname=%s password=%s ",
		DBHost, DBPort, DBUser, DBName, DBPassword)
	log.Printf("qqq Openning DB: %s ...\n", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("DB open failed:", connStr, err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("DB connection successful!!")
	return NewDBRepo(db)
	//return db
}
