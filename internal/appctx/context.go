package appctx

import (
	"apus-sample/common/database"
)

var Context *appContext

type appContext struct {
	DB *database.Database
}

func InitContext() error {
	var err error
	Context = &appContext{}
	Context.DB, err = database.New()
	return err
}
