package utils

import (
	"apus-sample/common/constant"
)
import "github.com/gin-gonic/gin"

func GetCompanyCode(c *gin.Context) string {
	companyCode := c.GetHeader("company_code")
	if companyCode == "" {
		companyCode = constant.DefaultSchema
	}
	return companyCode
}

func GetQueries(c *gin.Context) map[string]string {
	queries := make(map[string]string)
	_ = c.BindQuery(queries)
	return queries
}
