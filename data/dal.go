package data

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const insertData = `INSERT INTO map (id, val) VALUES ($1, $2)`
const readData = `SELECT id, val FROM map ORDER BY val DESC LIMIT $1;`

func (db *DB) UpdateMap(url string) {
	fmt.Println("Update Map from:", url)

	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()

	lines, err := LinesFromReader(resp.Body)
	checkError(err)

	for _, line := range lines {
		words := strings.Fields(line)
		_, err := db.Exec(insertData, words[0], words[1])
		checkError(err)
	}
}

func (db *DB) GetTopKth(k int) string {
	fmt.Printf("Get %v-th record \n", k)
	rows, err := db.Query(readData, k)
	checkError(err)
	defer rows.Close()

	var id string
	var val string
	for rows.Next() {
		err = rows.Scan(&id, &val)
		checkError(err)
	}
	return id
}

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
		os.Exit(1)
	}
}
