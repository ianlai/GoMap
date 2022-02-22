package data

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const insertData = `INSERT INTO map (uid, val) VALUES ($1, $2)`
const readData = `SELECT uid, val FROM map ORDER BY val DESC LIMIT $1`

func UpdateMap(db *sql.DB, url string) (err error) {
	fmt.Println("Update Map from:", url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	lines, err := LinesFromReader(resp.Body)
	checkError(err)

	for _, line := range lines {
		words := strings.Fields(line)
		_, err := db.Exec(insertData, words[0], words[1])
		checkError(err)
	}
	return
}

func GetTopKth(db *sql.DB, k int) string {
	fmt.Printf("Get %v-th record \n", k)
	rows, err := db.Query(readData, k)
	checkError(err)
	defer rows.Close()

	var id int
	var uid string
	var val string
	var updated_at time.Time
	for rows.Next() {
		err = rows.Scan(&id, &uid, &val, &updated_at)
		checkError(err)
	}
	return uid
}

// func (db *DB) UpdateMap(url string) {
// 	fmt.Println("Update Map from:", url)

// 	resp, err := http.Get(url)
// 	checkError(err)
// 	defer resp.Body.Close()

// 	lines, err := LinesFromReader(resp.Body)
// 	checkError(err)

// 	for _, line := range lines {
// 		words := strings.Fields(line)
// 		_, err := db.Exec(insertData, words[0], words[1])
// 		checkError(err)
// 	}
// }

// func (db *DB) GetTopKth(k int) string {
// 	fmt.Printf("Get %v-th record \n", k)
// 	rows, err := db.Query(readData, k)
// 	checkError(err)
// 	defer rows.Close()

// 	var id string
// 	var val string
// 	for rows.Next() {
// 		err = rows.Scan(&id, &val)
// 		checkError(err)
// 	}
// 	return id
// }

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
