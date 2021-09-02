package company

import (
	"apus-sample/common/utils"
	"apus-sample/internal/appctx"
	"context"
	"fmt"
)

func List(ctx context.Context) (companies []Company, err error) {
	conn := appctx.Context.DB.MustGet(ctx, utils.GetCompanyCodeFromCtx(ctx))
	fmt.Println(utils.GetCompanyCodeFromCtx(ctx))
	companies, err = listAllCompanies(conn)
	return
}
