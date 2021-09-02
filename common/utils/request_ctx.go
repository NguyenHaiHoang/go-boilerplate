package utils

import (
	"apus-sample/common/constant"
	"context"
)
import "github.com/gin-gonic/gin"

func GetCompanyCodeFromCtx(ctx context.Context) (companyCode string) {
	defer func() {
		if companyCode == ""{
			companyCode = constant.DefaultSchema
		}
	}()
	switch ctx.(type) {
	case *gin.Context:
		ginCtx := ctx.(*gin.Context)
		companyCode = ginCtx.GetHeader("company_code")
	default:
		companyCode = constant.DefaultSchema
	}
	return companyCode
}
