package rest

import "github.com/gin-gonic/gin"

func SetupRoute(router *gin.Engine) error {
	router.GET("/companies", ListCompanies)
	return nil
}
