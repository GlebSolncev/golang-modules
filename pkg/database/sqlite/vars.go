package sqlite

import (
	"context"
	"golang-modules/pkg/database/contracts"
	"golang-modules/pkg/ent"
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
