package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ianlai/GoMap/data"
)

const removedLength int64 = 500
const urlDefault string = "https://bucket-ian-1.s3.amazonaws.com/data_full.txt"
const numDefault int = 10

func main() {
	log.Println("Hello GoMap")
	var url string
	var num int
	flag.StringVar(&url, "url", urlDefault, "The URL to download the data")
	flag.IntVar(&num, "num", numDefault, "Show the largest number records based on value")
	flag.Parse()
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("NUM: %v\n", num)

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

	for i, record := range records {
		//log.Printf("Top-%vth -> Uid: %v, Val:%v\n", k, record.Uid, record.Val)
		fmt.Printf("%v: %v\n", i, record.Uid)
	}
}
