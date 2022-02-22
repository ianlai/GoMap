package main

import (
	"fmt"

	"github.com/ianlai/GoMap/data"
)

func main() {
	fmt.Println("Hello Map")
	url := "https://bucket-ian-1.s3.amazonaws.com/data_short.txt"
	k := 29

	db := data.InitDB()
	data.UpdateMap(db, url)
	res := data.GetTopKth(db, k)
	fmt.Printf("Top-%vth: %v\n", k, res)
}
