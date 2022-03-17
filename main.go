package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ianlai/GoMap/app"
	"github.com/ianlai/GoMap/data"
)

const removedLength int64 = 500
const urlDefault string = "https://bucket-ian-1.s3.amazonaws.com/data_full.txt"
const numDefault int = 10

func main() {
	log.Println("Hello GoMap")

	//Read flag
	var url string
	var num int
	flag.StringVar(&url, "url", urlDefault, "The URL to download the data")
	flag.IntVar(&num, "num", numDefault, "Show the given number of records which have largest values")
	flag.Parse()
	log.Printf("Read flag - url: %s\n", url)
	log.Printf("Read flag - num: %v\n", num)

	//Set database
	db := data.InitDB()
	log.Println("DB initialized.")

	//Set server
	server := &app.Server{
		Repo: db,
		Name: "GoMap",
	}
	log.Println("Server created.")

	//Set router
	r := chi.NewRouter()
	server.SetRouter(r)
	log.Println("Router set.")

	//Run importer
	lines, err := RetrieveData(url, removedLength)
	if err != nil {
		log.Printf("%s", err)
	}

	err = InsertRecords(db, lines)
	if err != nil {
		log.Printf("%s", err)
	}

	//Show top-k results in log
	records, err := GetTopKRecords(db, num)
	if err != nil {
		log.Printf("%s", err)
	}

	for _, record := range records {
		log.Printf("%v\n", record.Uid)
	}

	//Start server
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
