package contracts

type (
	DBMethods interface {
		Get() []byte
		Save(data []byte) bool
	}
)
