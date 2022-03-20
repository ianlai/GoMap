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
	DBPort     = "5432"

	//(3a) Connect from Docker through Internal IP
	//DBHost = "172.27.0.2"

	//(3b) Connect from host machine through Host IP
	//DBHost = "localhost" //connect by host

	//(3c) Connect by docker-compose
	//DBHost = "db"
)

// Repo interface
type Repo interface {
	InsertRecord(string, string) (int64, error)
	ListRecords(int, bool) ([]Record, error)
	GetRecordByUid(string) (*Record, error)
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
	//Get the DBHOST env variable to connect to database
	DBHost := os.Getenv("DBHOST")

	connStr := fmt.Sprintf(
		"sslmode=disable host=%s port=%s user=%s dbname=%s password=%s ",
		DBHost, DBPort, DBUser, DBName, DBPassword)
	log.Printf("Openning DB: %s \n", connStr)

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
}
