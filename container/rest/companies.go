package rest

import (
	"apus-sample/common/utils"
	"apus-sample/services/company"
	"github.com/gin-gonic/gin"
)

func ListCompanies(c *gin.Context) {
	companies, err := company.List(c, utils.GetCompanyCode(c), utils.GetQueries(c))
	utils.PanicWhenError(err)
	c.JSON(200, companies)
}
