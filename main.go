package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/router"
)

var (
	controllerRegister router.RouterRegisterInterface = router.NewRouterRegister()
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1/fitness")
	controllerRegister.TrackingRouter(v1.Group("/tracking"))
	controllerRegister.GoalsRouter(v1.Group("/goals"))
	r.Run(":8081")
}
