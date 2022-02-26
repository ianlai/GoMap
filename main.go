package main

import (
	"fmt"
	"log"

	"github.com/ianlai/GoMap/data"
)

const removedLength int64 = 500

func main() {
	log.Println("Hello GoMap")
	//url := "https://bucket-ian-1.s3.amazonaws.com/data_prefix.txt"
	k := 13
	url := "https://bucket-ian-1.s3.amazonaws.com/data_full.txt"

	db := data.InitDB()

	lines, err := RetrieveData(url, removedLength)
	if err != nil {
		log.Printf("%s", err)
	}

	err = InsertRecords(db, lines)
	if err != nil {
		log.Printf("%s", err)
	}

	records, err := GetTopKRecords(db, k)
	if err != nil {
		log.Printf("%s", err)
	}

	for i, record := range records {
		//log.Printf("Top-%vth -> Uid: %v, Val:%v\n", k, record.Uid, record.Val)
		fmt.Printf("%v: %v\n", i, record.Uid)
	}
}
