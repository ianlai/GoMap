package main

import (
	"fmt"

	"github.com/ianlai/GoMap/data"
)

func main() {
	fmt.Println("Hello Map")
	//url := "https://bucket-ian-1.s3.amazonaws.com/data_short.txt"
	k := 29

	db := data.InitDB()
	//db.UpdateMap(url)
	res := db.GetTopKth(k)
	fmt.Printf("Top-%vth: %v\n", k, res)
}
