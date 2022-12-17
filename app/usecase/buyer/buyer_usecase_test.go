package buyer

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"gokomodo-assesment/app/domain/entity"

	"github.com/gin-gonic/gin"
	"gokomodo-assesment/app/domain/model"
	"gokomodo-assesment/app/helper/mock/repository"
	"testing"
)

func TestBuyerUsecase_Login(t *testing.T) {

	var c *gin.Context
	mockRepoBuyer := &repository.MockBuyerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	request := &model.LoginRequestModel{}
	buyerEntity := &entity.Buyer{}
	//response := &model.LoginResponseModel{}

	Convey("Test Usecase Buyer Login", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When Login Success, should return data", func() {
				mockRepoBuyer.On("GetBuyer", buyerEntity).Return(buyerEntity, nil).Once()
				mockRepoJWT.On("CheckPassword", "", "").Return(nil).Once()
				mockRepoJWT.On("GenerateJWT", "", "buyer").Return("token", nil).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				res, err := uc.Login(c, request)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When get buyer failed, should return error", func() {
				mockRepoBuyer.On("GetBuyer", buyerEntity).Return(nil, errors.New("errors")).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				res, err := uc.Login(c, request)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When password invalid, should return error", func() {
				mockRepoBuyer.On("GetBuyer", buyerEntity).Return(buyerEntity, nil).Once()
				mockRepoJWT.On("CheckPassword", "", "").Return(errors.New("error")).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				res, err := uc.Login(c, request)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When generate jwt failed, should return error", func() {
				mockRepoBuyer.On("GetBuyer", buyerEntity).Return(buyerEntity, nil).Once()
				mockRepoJWT.On("CheckPassword", "", "").Return(nil).Once()
				mockRepoJWT.On("GenerateJWT", "", "buyer").Return("", errors.New("error")).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				res, err := uc.Login(c, request)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestBuyerUsecase_GetListProduct(t *testing.T) {

	var c *gin.Context
	mockRepoBuyer := &repository.MockBuyerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	productEntity := []*entity.Product{}
	productEntity = append(productEntity, &entity.Product{ProductName: "tes"})
	//response := &model.LoginResponseModel{}

	Convey("Test Usecase Buyer GetListProduct", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When get list product Success, should return data", func() {
				mockRepoProduct.On("GetListProduct").Return(productEntity, nil).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				res, err := uc.GetListProduct(c)
				So(res, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When get list product failed, should return error", func() {
				mockRepoProduct.On("GetListProduct").Return(nil, errors.New("errors")).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				res, err := uc.GetListProduct(c)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestBuyerUsecase_AddOrder(t *testing.T) {

	var c *gin.Context
	mockRepoBuyer := &repository.MockBuyerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	orderEntity := &entity.Order{Status: "Pending"}

	Convey("Test Usecase Buyer AddOrder", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When add order Success, should not return error", func() {
				mockRepoOrder.On("AddOrder", orderEntity).Return(nil).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				err := uc.AddOrder(c, &model.AddOrderModel{})
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When add order failed, should return error", func() {
				mockRepoOrder.On("AddOrder", orderEntity).Return(errors.New("errors")).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				err := uc.AddOrder(c, &model.AddOrderModel{})
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestBuyerUsecase_GetListOrder(t *testing.T) {

	var c *gin.Context
	mockRepoBuyer := &repository.MockBuyerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	orderEntity := []*entity.Order{}
	orderEntity = append(orderEntity, &entity.Order{Price: 1})

	Convey("Test Usecase Buyer GetListOrder", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When get list order Success, should return data", func() {
				mockRepoOrder.On("GetOrderByBuyerID", "").Return(orderEntity, nil).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				res, err := uc.GetListOrder(c, "")
				So(res, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When get list order failed, should return error", func() {
				mockRepoOrder.On("GetOrderByBuyerID", "").Return(nil, errors.New("errors")).Once()
				uc := NewBuyerUsecase(mockRepoJWT, mockRepoBuyer, mockRepoProduct, mockRepoOrder)
				res, err := uc.GetListOrder(c, "")
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
