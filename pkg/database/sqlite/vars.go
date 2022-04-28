package sqlite

import (
	"context"
	"crud/pkg/database/contracts"
	"crud/pkg/ent"
)

type (
	Method struct {
		contracts.DBMethods
	}
)

var (
	Client         *ent.Client
	ctx            context.Context
	DataSourceName string //Example: "file:./test.db?mode=rwc&cache=shared&_fk=1"
)
