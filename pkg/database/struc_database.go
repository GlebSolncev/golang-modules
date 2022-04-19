package database

type (
	StructMethod struct {
		Method DBMethods
	}

	database struct {
		Data []byte
	}
)

var (
	db = database{}
)

func (m *StructMethod) Get() []byte {
	return db.Data
}
func (m *StructMethod) Save(data []byte) bool {
	db.Data = data
	return true
}
