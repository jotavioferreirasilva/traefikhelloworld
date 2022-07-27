package routes

import (
	"helloworld/controllers"
	"net/http"
)

var helloworld = []Route{
	{
		URI:      "/helloworld",
		Method:   http.MethodGet,
		Function: controllers.HelloWorld,
	},
}
