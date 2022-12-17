package router

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assesment/app/controller"
	"gokomodo-assesment/app/infrastructure/middleware"
)

func BuyerRouter(c *gin.RouterGroup, buyerService *controller.BuyerController) {
	c.POST("/login", buyerService.Login)
	c.GET("/products", middleware.AuthenticationBuyerMiddleware, buyerService.GetListProduct)
	c.GET("/orders", middleware.AuthenticationBuyerMiddleware, buyerService.GetListOrder)
	c.POST("/order", middleware.AuthenticationBuyerMiddleware, buyerService.AddOrder)
}
