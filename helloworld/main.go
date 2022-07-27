package main

import (
	"fmt"
	"helloworld/router"
)

func main() {
	router := router.Generate()
	router.Run(fmt.Sprintf(":%d", 5000))
}
