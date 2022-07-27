package routes

import "github.com/gin-gonic/gin"

type Route struct {
	URI      string
	Method   string
	Function func(c *gin.Context)
}

func Configure(router *gin.Engine) *gin.Engine {
	var routes []Route

	routes = append(routes, helloworld...)

	for _, route := range routes {
		router.Handle(route.Method, route.URI, route.Function)
	}

	return router
}
