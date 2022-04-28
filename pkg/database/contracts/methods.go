package contracts

type SaveType string

const (
	New    SaveType = "New"
	Update SaveType = "Update"
	Delete SaveType = "Delete"
)

type (
	DBMethods interface {
		Get() []byte
		Save(data []byte, sType SaveType) bool
	}
)
