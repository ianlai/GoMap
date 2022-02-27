package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ianlai/GoMap/data"
)

const removedLength int64 = 500
const urlDefault string = "https://amp-technical-challenge.s3.ap-northeast-1.amazonaws.com/sw-engineer-challenge.txt"
const numDefault int = 10

func main() {
	log.Println("Hello GoMap")

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
}
