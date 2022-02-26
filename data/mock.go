package data

type MockDB struct {
	MockInsertRecord          func(string, string) error
	MockGetRecordsSortedByVal func(int) ([]Record, error)
}

func (m *MockDB) InsertRecord(uid string, val string) error {
	return m.MockInsertRecord(uid, val)
}

func (m *MockDB) GetRecordsSortedByVal(k int) ([]Record, error) {
	return m.MockGetRecordsSortedByVal(k)
}
