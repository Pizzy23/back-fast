package routers

import (
	handlers "fast/internal"

	"github.com/gin-gonic/gin"
)

func RegisterShareDataRoutes(router *gin.Engine) {
	shareDataGroup := router.Group("/share-data")
	{
		shareDataGroup.GET("/", handlers.GetShareData)
		shareDataGroup.POST("/", handlers.PostShareData)
		shareDataGroup.PUT("/", handlers.PutShareData)
	}
}

func RegisterDataStandardizedRoutes(router *gin.Engine) {
	dataStandardizedGroup := router.Group("/data-standardized")
	{
		dataStandardizedGroup.GET("/", handlers.GetDataStandardized)
		dataStandardizedGroup.POST("/", handlers.PostDataStandardized)
		dataStandardizedGroup.PUT("/", handlers.PutDataStandardized)
	}
}
