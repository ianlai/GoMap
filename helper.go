package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ianlai/GoMap/data"
)

func RetrieveData(url string, removeLength int64) ([]string, error) {
	log.Println("[Main] RetrieveData from:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r, err := RemovePrefixData(resp.Body, removeLength)
	if err != nil {
		return nil, err
	}
	lines, err := GetLinesFromReader(r)
	if err != nil {
		return nil, err
	}
	return lines, nil
}
func RemovePrefixData(r io.Reader, removeLength int64) (io.Reader, error) {
	_, err := io.CopyN(ioutil.Discard, r, removeLength)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func GetLinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	count := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		fmt.Println(count, lines)
		count++
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}
func InsertRecords(db *sql.DB, records []string) error {
	log.Println("[Main] InsertRecords")
	for _, record := range records {
		words := strings.Fields(record)
		err := data.InsertRecord(db, words[0], words[1])
		if err != nil {
			return err
		}
	}
	return nil
}
func GetTopKthVal(db *sql.DB, k int) (string, error) {
	log.Printf("[Main] Get %v-th record \n", k)
	records, err := data.GetRecordsSortedByVal(db, k)
	if err != nil {
		return "", err
	}
	if len(records) < k {
		return "", errors.New("k is larger than the size of the data")
	}
	return records[len(records)-1].Uid, nil
}
