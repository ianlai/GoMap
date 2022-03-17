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
func (db *DB) GetRecordByUid(uid string) (*Record, error) {
	query := `
		SELECT 
			* 
		FROM 
			map 
		WHERE
			uid = $1
			`
	log.Printf("[DAL] GetRecordByUid: %s", uid)
	row := db.QueryRow(query, uid)

	r := &Record{}
	err := row.Scan(&r.ID, &r.Uid, &r.Val)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (db *DB) ListRecords(k int, isSortedByVal bool) ([]Record, error) {
	var query string
	if isSortedByVal {
		query = `
		SELECT 
			* 
		FROM 
			map 
		ORDER BY 
			val DESC 
		LIMIT $1`
	} else {
		query = `
		SELECT 
			* 
		FROM 
			map 
		LIMIT $1`
	}
	log.Printf("[DAL] ListRecords: %v", k)
	rows, err := db.Query(query, k)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var records []Record
	for rows.Next() {
		var r Record
		err = rows.Scan(&r.ID, &r.Uid, &r.Val)
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, nil
}
