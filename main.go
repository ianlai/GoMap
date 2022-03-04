package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ianlai/GoMap/app"
	"github.com/ianlai/GoMap/data"
)

const removedLength int64 = 500
const urlDefault string = "https://bucket-ian-1.s3.amazonaws.com/data_full.txt"
const numDefault int = 10

type UserInfo struct {
	Name string `json:"namename"`
	Age  int    `json:"ageage"`
}

func main() {
	log.Println("Hello GoMap")

	//Set routes
	log.Println("Server started..")

	r := chi.NewRouter()
	app.SetRouter(r)

	//Read flag
	var url string
	var num int
	flag.StringVar(&url, "url", urlDefault, "The URL to download the data")
	flag.IntVar(&num, "num", numDefault, "Show the given number of records which have largest values")
	flag.Parse()
	log.Printf("url: %s\n", url)
	log.Printf("num: %v\n", num)

	db := data.InitDB()

	lines, err := RetrieveData(url, removedLength)
	if err != nil {
		log.Printf("%s", err)
	}

	err = InsertRecords(db, lines)
	if err != nil {
		log.Printf("%s", err)
	}

	records, err := GetTopKRecords(db, num)
	if err != nil {
		log.Printf("%s", err)
	}

	//Show the final result in the requested format
	fmt.Println("========== Final Result (The IDs of the largest values) ==========")
	for _, record := range records {
		fmt.Printf("%v\n", record.Uid)
	}

	//Start server
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
