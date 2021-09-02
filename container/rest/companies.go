package rest

import (
	"apus-sample/common/utils"
	"apus-sample/services/company"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListCompanies(c *gin.Context) {
	fmt.Println("debug")
	companies, err := company.List(c)
	utils.PanicWhenError(err)
	c.JSON(200, companies)
}
