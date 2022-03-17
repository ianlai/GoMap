package data

type MockDB struct {
	MockInsertRecord   func(string, string) error
	MockListRecords    func(int, bool) ([]Record, error)
	MockGetRecordByUid func(string) (*Record, error)
}

func (m *MockDB) InsertRecord(uid string, val string) error {
	return m.MockInsertRecord(uid, val)
}

func (m *MockDB) GetRecordByUid(uid string) (*Record, error) {
	return m.MockGetRecordByUid(uid)
}

func (m *MockDB) ListRecords(k int, isSortedByVal bool) ([]Record, error) {
	return m.MockListRecords(k, isSortedByVal)
}
