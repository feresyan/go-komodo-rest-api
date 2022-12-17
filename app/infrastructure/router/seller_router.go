package router

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assesment/app/controller"
	"gokomodo-assesment/app/infrastructure/middleware"
)

func SellerRouter(c *gin.RouterGroup, sellerService *controller.SellerController) {
	c.POST("/login", sellerService.Login)
	c.GET("/products", middleware.AuthenticationSellerMiddleware, sellerService.GetListProduct)
	c.POST("/products", middleware.AuthenticationSellerMiddleware, sellerService.AddProduct)
	c.PUT("/orders", middleware.AuthenticationSellerMiddleware, sellerService.AcceptOrder)
	c.GET("/orders", middleware.AuthenticationSellerMiddleware, sellerService.GetListOrder)
}
