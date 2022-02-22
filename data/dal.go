package data

import (
	"database/sql"
	"log"
)

const insertData = `
	INSERT INTO 
		map (uid, val) 
	VALUES 
		($1, $2)`

const readData = `
	SELECT 
		uid, val 
	FROM 
		map 
	ORDER BY 
		val 
	DESC LIMIT $1`

func InsertRecord(db *sql.DB, uid string, val string) error {
	log.Printf("[Data] InsertRecord: %s, %s", uid, val)
	_, err := db.Exec(insertData, uid, val)
	if err != nil {
		return err
	}
	return nil
}

func GetRecordsSortedByVal(db *sql.DB, k int) ([]Record, error) {
	log.Printf("[Data] GetRecordsSortedByVal: %v", k)

	rows, err := db.Query(readData, k)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var records []Record

	for rows.Next() {
		var r Record
		err = rows.Scan(&r.Uid, &r.Val)
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, nil
}
