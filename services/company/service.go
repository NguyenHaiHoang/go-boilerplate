package company

import (
	"apus-sample/internal/appctx"
	"context"
)

func List(ctx context.Context, companyCode string, filters map[string]string) (companies []Company, err error) {
	conn := appctx.Context.DB.MustGet(ctx, companyCode)
	companies, err = listCompanies(conn, filters)
	return
}
