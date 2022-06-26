package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/controller"
)

var (
	controllerRegister controller.ControllerRegisterInterface = controller.NewController()
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	controllerRegister.TrackingController(v1.Group("/tracking"))
	r.Run()
}
