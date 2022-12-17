package controller

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assesment/app/domain/model"
	"gokomodo-assesment/app/usecase/buyer"
	"net/http"
)

type BuyerController struct {
	uc buyer.BuyerPort
}

func NewBuyerController(u buyer.BuyerPort) *BuyerController {
	return &BuyerController{
		uc: u,
	}
}

type BuyerIDBindingObject struct {
	BuyerID string `json:"buyer_id" binding:"required"`
}

func (s *BuyerController) Login(c *gin.Context) {
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

	c.JSON(200, res)
	return
}

func (s *BuyerController) GetListProduct(c *gin.Context) {
	res, err := s.uc.GetListProduct(c.Copy())
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (s *BuyerController) AddOrder(c *gin.Context) {
	var order *model.AddOrderModel
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusNotAcceptable, "invalid object")
		return
	}

	err := s.uc.AddOrder(c.Copy(), order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, &model.OrderResponseModel{
		ResponseCode: "00",
		ResponseDesc: "Add Order Success",
	})
	return
}

func (s *BuyerController) GetListOrder(c *gin.Context) {
	var buyer *BuyerIDBindingObject
	if err := c.BindJSON(&buyer); err != nil {
		c.JSON(http.StatusBadRequest, "buyer id invalid")
		return
	}

	res, err := s.uc.GetListOrder(c.Copy(), buyer.BuyerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
