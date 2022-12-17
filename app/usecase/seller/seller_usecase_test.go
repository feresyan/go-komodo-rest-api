package seller

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/gin-gonic/gin"
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/domain/model"
	"gokomodo-assesment/app/helper/mock/repository"
	"testing"
)

func TestSellerUsecase_Login(t *testing.T) {

	var c *gin.Context
	mockRepoSeller := &repository.MockSellerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	request := &model.LoginRequestModel{}
	sellerEntity := &entity.Seller{}

	Convey("Test Usecase Seller Login", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When Login Success, should return data", func() {
				mockRepoSeller.On("GetSeller", sellerEntity).Return(sellerEntity, nil).Once()
				mockRepoJWT.On("CheckPassword", "", "").Return(nil).Once()
				mockRepoJWT.On("GenerateJWT", "", "seller").Return("token", nil).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				res, err := uc.Login(c, request)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When get seller failed, should return error", func() {
				mockRepoSeller.On("GetSeller", sellerEntity).Return(nil, errors.New("errors")).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				res, err := uc.Login(c, request)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When password invalid, should return error", func() {
				mockRepoSeller.On("GetSeller", sellerEntity).Return(sellerEntity, nil).Once()
				mockRepoJWT.On("CheckPassword", "", "").Return(errors.New("error")).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				res, err := uc.Login(c, request)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When generate jwt failed, should return error", func() {
				mockRepoSeller.On("GetSeller", sellerEntity).Return(sellerEntity, nil).Once()
				mockRepoJWT.On("CheckPassword", "", "").Return(nil).Once()
				mockRepoJWT.On("GenerateJWT", "", "seller").Return("", errors.New("error")).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				res, err := uc.Login(c, request)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSellerUsecase_GetListProduct(t *testing.T) {

	var c *gin.Context
	mockRepoSeller := &repository.MockSellerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	var productEntity []*entity.Product
	productEntity = append(productEntity, &entity.Product{ProductName: "tes"})

	Convey("Test Usecase Seller GetListProduct", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When get list product Success, should return data", func() {
				mockRepoProduct.On("GetListProductBySellerID", "").Return(productEntity, nil).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				res, err := uc.GetListProduct(c, "")
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When get list product failed, should return error", func() {
				mockRepoProduct.On("GetListProductBySellerID", "").Return(nil, errors.New("errors")).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				res, err := uc.GetListProduct(c, "")
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSellerUsecase_AddNewProduct(t *testing.T) {
	var c *gin.Context
	mockRepoSeller := &repository.MockSellerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	Convey("Test Usecase Seller AddNewProduct", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When add new product Success, should return success", func() {
				mockRepoProduct.On("AddProduct", &entity.Product{}).Return(nil).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				err := uc.AddNewProduct(c, &model.ProductModel{})
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When add new product failed, should return error", func() {
				mockRepoProduct.On("AddProduct", &entity.Product{}).Return(errors.New("errors")).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				err := uc.AddNewProduct(c, &model.ProductModel{})
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSellerUsecase_GetListOrder(t *testing.T) {
	var c *gin.Context
	mockRepoSeller := &repository.MockSellerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	orderEntity := []*entity.Order{&entity.Order{SellerID: "test"}}

	Convey("Test Usecase Seller GetListOrder", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When get list order by seller id Success, should return data", func() {
				mockRepoOrder.On("GetOrderBySellerID", "").Return(orderEntity, nil).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				res, err := uc.GetListOrder(c, "")
				So(res, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When add new product failed, should return error", func() {
				mockRepoOrder.On("GetOrderBySellerID", "").Return(nil, errors.New("errors")).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				res, err := uc.GetListOrder(c, "")
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSellerUsecase_AcceptOrder(t *testing.T) {
	var c *gin.Context
	mockRepoSeller := &repository.MockSellerRepository{}
	mockRepoJWT := &repository.MockJWTRepository{}
	mockRepoProduct := &repository.MockProductRepository{}
	mockRepoOrder := &repository.MockOrderRepository{}

	Convey("Test Usecase Seller AcceptOrder", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When accept order Success, should return success", func() {
				mockRepoOrder.On("AcceptOrder", &entity.Order{}).Return(nil).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				err := uc.AcceptOrder(c, &model.AcceptOrderModel{})
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When accept order failed, should return error", func() {
				mockRepoOrder.On("AcceptOrder", &entity.Order{}).Return(errors.New("errors")).Once()
				uc := NewSellerUsecase(mockRepoJWT, mockRepoSeller, mockRepoProduct, mockRepoOrder)
				err := uc.AcceptOrder(c, &model.AcceptOrderModel{})
				So(err, ShouldNotBeNil)
			})
		})
	})
}
