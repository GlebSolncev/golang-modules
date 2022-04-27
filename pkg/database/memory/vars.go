package memory

import (
	"crud/pkg/database/contracts"
)

type (
	Method struct {
		contracts.DBMethods
	}

	database struct {
		Data []byte
	}
)

var (
	db = database{}
)
