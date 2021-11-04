//Starter HTTP Server with Gin Framework
package main

import (
	"kstarter/service/controller"
	"kstarter/service/routes"
)

func main() {
	// Creates a gin router with default middleware:
	// Controller handlers are added here
	var controller controller.Controller = &controller.DefaultController{}
	//setup routes
	r := routes.SetupRouter(controller)

	// running
	r.Run()
}
