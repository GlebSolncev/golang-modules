package database

type (
	DBMethods interface {
		Get() []byte
		Save(data []byte) bool
	}

	NewMethod struct{}
)

var (
	defaultMethod DBMethods = &FileMethod{}
)

func (n NewMethod) Start() DBMethods {
	return defaultMethod
}

var filename = "storage/todo.json"

func check(err error) {
	if err != nil {
		panic(err)
	}
}
