package data

import (
	"log"
)

func (db *DB) InsertRecord(uid string, val string) error {
	const insertData = `
		INSERT INTO 
			map (uid, val) 
		VALUES 
			($1, $2)`

	log.Printf("[DAL] InsertRecord: %s, %s", uid, val)
	_, err := db.Exec(insertData, uid, val)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetRecordsSortedByVal(k int) ([]Record, error) {
	const readData = `
		SELECT 
			uid, val 
		FROM 
			map 
		ORDER BY 
			val 
		DESC 
		LIMIT $1`

	log.Printf("[DAL] GetRecordsSortedByVal: %v", k)
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
