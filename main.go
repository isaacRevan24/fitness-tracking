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
	v1 := r.Group("/v1")
	controllerRegister.TrackingRouter(v1.Group("/tracking"))
	r.Run()
}
