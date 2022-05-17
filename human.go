package golang_modules

//go:generate go run ./cmd/structuring/main.go -name Human

type (
	Human struct {
		FirstName string
		LastName  string
		Age       int
	}
)
