package controllers

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	c.Writer.WriteString("Hello World!")
}
