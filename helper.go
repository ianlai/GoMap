package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ianlai/GoMap/data"
)

func RetrieveData(url string, removeLength int64) ([]string, error) {
	log.Println("[Helper] RetrieveData from:", url)
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
		count++
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}
func InsertRecords(repo data.Repo, records []string) error {
	log.Println("[Helper] InsertRecords")
	for _, record := range records {
		words := strings.Fields(record)
		err := repo.InsertRecord(words[0], words[1])
		if err != nil {
			return err
		}
	}
	return nil
}
func GetTopKRecords(repo data.Repo, k int) ([]data.Record, error) {
	log.Printf("[Helper] Get top-%v records \n", k)
	records, err := repo.ListRecords(k, true)
	if err != nil {
		return []data.Record{}, err
	}
	if len(records) < k {
		return []data.Record{}, errors.New("k is larger than the size of the data")
	}
	return records, nil
}
