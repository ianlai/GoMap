package main

import (
	"log"

	"github.com/ianlai/GoMap/data"
)

func main() {
	log.Println("Hello Map")
	url := "https://bucket-ian-1.s3.amazonaws.com/data_prefix.txt"
	k := 29

	db := data.InitDB()

	lines, err := RetrieveData(db, url)
	if err != nil {
		log.Printf("%s", err)
	}

	err = InsertRecords(db, lines)
	if err != nil {
		log.Printf("%s", err)
	}

	uid, err := GetTopKthVal(db, k)
	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("[Final] Top-%vth: %v\n", k, uid)
}
