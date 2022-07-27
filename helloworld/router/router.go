package router

import (
	"github.com/gin-gonic/gin"
	"helloworld/router/routes"
)

func Generate() *gin.Engine {
	router := gin.Default()
	return routes.Configure(router)
}
