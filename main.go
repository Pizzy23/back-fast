package main

import (
	"fast/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.RegisterShareDataRoutes(router)
	routers.RegisterDataStandardizedRoutes(router)

	router.Run(":8080")
}
