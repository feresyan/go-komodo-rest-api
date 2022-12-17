package controller

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assesment/app/domain/model"
	"gokomodo-assesment/app/usecase/seller"
	"net/http"
)

type SellerController struct {
	uc seller.SellerPort
}

type SellerIDBindingObject struct {
	SellerID string `json:"seller_id"`
}

func NewSellerController(u seller.SellerPort) *SellerController {
	return &SellerController{
		uc: u,
	}
}

func (s *SellerController) Login(c *gin.Context) {
	var loginRequest *model.LoginRequestModel
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusNotAcceptable, "invalid object")
		return
	}

	res, err := s.uc.Login(c.Copy(), loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Set token to cookie ( optional )
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", res.Token, 3600*3, "", "", false, true)

	c.JSON(http.StatusOK, res)
	return
}

func (s *SellerController) GetListProduct(c *gin.Context) {
	var seller *SellerIDBindingObject
	if err := c.BindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, "seller id invalid")
		return
	}

	res, err := s.uc.GetListProduct(c.Copy(), seller.SellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (s *SellerController) AddProduct(c *gin.Context) {
	var product *model.ProductModel
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusNotAcceptable, "invalid object")
		return
	}

	err := s.uc.AddNewProduct(c.Copy(), product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, &model.ProductResponseModel{
		ResponseCode: "00",
		ResponseDesc: "Add Product Success",
	})
	return
}

func (s *SellerController) GetListOrder(c *gin.Context) {
	var seller *SellerIDBindingObject
	if err := c.BindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, "seller id invalid")
		return
	}

	res, err := s.uc.GetListOrder(c.Copy(), seller.SellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (s *SellerController) AcceptOrder(c *gin.Context) {
	var request *model.AcceptOrderModel
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusNotAcceptable, "invalid object")
		return
	}

	err := s.uc.AcceptOrder(c.Copy(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, &model.OrderResponseModel{
		ResponseCode: "00",
		ResponseDesc: "Accept Order Success",
	})
	return
}
