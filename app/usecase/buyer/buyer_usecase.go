package buyer

import (
	"context"
	"errors"
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/domain/model"
	"gokomodo-assesment/app/domain/repository"
)

type BuyerUsecase struct {
	repoJWT     repository.JWTRepository
	repoBuyer   repository.BuyerRepository
	repoProduct repository.ProductRepository
	repoOrder   repository.OrderRepository
}

func NewBuyerUsecase(rj repository.JWTRepository, rb repository.BuyerRepository, rp repository.ProductRepository, ro repository.OrderRepository) BuyerPort {
	return &BuyerUsecase{
		repoJWT:     rj,
		repoBuyer:   rb,
		repoProduct: rp,
		repoOrder:   ro,
	}
}

func (u *BuyerUsecase) Login(ctx context.Context, request *model.LoginRequestModel) (*model.LoginResponseModel, error) {

	// Get buyer from env
	getBuyer := &entity.Buyer{
		Email:    request.Email,
		Password: "",
	}

	buyer, err := u.repoBuyer.GetBuyer(getBuyer)
	if err != nil {
		return nil, errors.New("invalid Email or Password")
	}

	// check if password valid
	err = u.repoJWT.CheckPassword(buyer.Password, request.Password)
	if err != nil {
		return nil, errors.New("invalid Email or Password")
	}

	// generate jwt
	token, err := u.repoJWT.GenerateJWT(request.Email, "buyer")
	if err != nil {
		return nil, errors.New("failed to create token")
	}

	//set response
	response := &model.LoginResponseModel{
		Token: token,
	}

	// return res
	return response, nil
}

func (u *BuyerUsecase) GetListProduct(ctx context.Context) ([]*model.ProductModel, error) {

	var response []*model.ProductModel

	listProduct, err := u.repoProduct.GetListProduct()
	if err != nil {
		return nil, errors.New("get list product failed")
	}

	for i := 0; i < len(listProduct); i++ {
		response = append(response, &model.ProductModel{
			ID:          listProduct[i].ID,
			SellerID:    listProduct[i].SellerID,
			ProductName: listProduct[i].ProductName,
			Description: listProduct[i].Description,
			Price:       listProduct[i].Price,
		})
	}

	// return response
	return response, nil
}

func (u *BuyerUsecase) AddOrder(ctx context.Context, request *model.AddOrderModel) error {
	order := &entity.Order{
		SellerID:                   request.SellerID,
		BuyerID:                    request.BuyerID,
		ProductID:                  request.ProductID,
		DeliverySourceAddress:      request.DeliverySourceAddress,
		DeliveryDestinationAddress: request.DeliveryDestinationAddress,
		Quantity:                   request.Quantity,
		Price:                      request.Price,
		TotalPrice:                 request.TotalPrice,
		Status:                     "Pending",
	}

	err := u.repoOrder.AddOrder(order)
	if err != nil {
		return errors.New("failed to add order")
	}

	return nil
}

func (u *BuyerUsecase) GetListOrder(ctx context.Context, buyerID string) ([]*model.BuyerOrderModel, error) {
	var response []*model.BuyerOrderModel
	listOrder, err := u.repoOrder.GetOrderByBuyerID(buyerID)
	if err != nil {
		return nil, errors.New("failed to get list order")
	}

	for _, item := range listOrder {
		response = append(response, &model.BuyerOrderModel{
			ID:                         item.ID,
			BuyerID:                    item.BuyerID,
			BuyerName:                  item.Buyer.Name,
			SellerID:                   item.SellerID,
			SellerName:                 item.Seller.Name,
			DeliverySourceAddress:      item.DeliverySourceAddress,
			DeliveryDestinationAddress: item.DeliveryDestinationAddress,
			ProductID:                  item.ProductID,
			ProductName:                item.Product.ProductName,
			Quantity:                   item.Quantity,
			Price:                      item.Price,
			TotalPrice:                 item.TotalPrice,
			Status:                     item.Status,
		})
	}

	return response, nil
}
