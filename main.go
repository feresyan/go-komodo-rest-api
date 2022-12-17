package main

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assesment/app/adapter"
	"gokomodo-assesment/app/controller"
	"gokomodo-assesment/app/infrastructure/db_connection"
	"gokomodo-assesment/app/infrastructure/router"
	"gokomodo-assesment/app/usecase/buyer"
	"gokomodo-assesment/app/usecase/seller"
)

func main() {
	route := gin.Default()

	// Set default router
	defaultRouter := route.Group("/api/v1")

	// Set all repository
	jwtRepository := adapter.NewJWTAdapter()
	buyerRepository := adapter.NewBuyerAdapter(db_connection.DbConnection)
	sellerRepository := adapter.NewSellerAdapter(db_connection.DbConnection)
	productRepository := adapter.NewProductAdapter(db_connection.DbConnection)
	orderRepository := adapter.NewOrderAdapter(db_connection.DbConnection)

	// Set All Usecase
	buyerUsecase := buyer.NewBuyerUsecase(jwtRepository, buyerRepository, productRepository, orderRepository)
	sellerUsecase := seller.NewSellerUsecase(jwtRepository, sellerRepository, productRepository, orderRepository)

	// Set All Controller
	buyerController := controller.NewBuyerController(buyerUsecase)
	sellerController := controller.NewSellerController(sellerUsecase)

	// Set Router for Buyer
	routerBuyer := defaultRouter.Group("/buyer")
	router.BuyerRouter(routerBuyer, buyerController)

	// Set Router for Seller
	routerSeller := defaultRouter.Group("/seller")
	router.SellerRouter(routerSeller, sellerController)

	route.Run(":8801")
}
