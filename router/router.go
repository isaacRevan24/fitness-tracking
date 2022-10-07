//go:generate go run github.com/golang/mock/mockgen -source router.go -destination mock/router_mock.go -package mock
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/model"
)

var (
	mapper model.FitnessMapper = model.NewFitnessMapper()
)

type routerRegister struct{}

type RouterRegisterInterface interface {
	TrackingRouter(router *gin.RouterGroup)
	GoalsRouter(router *gin.RouterGroup)
}

func NewRouterRegister() RouterRegisterInterface {
	return &routerRegister{}
}
