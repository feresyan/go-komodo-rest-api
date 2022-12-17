package seller

import (
	"context"
	"errors"
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/domain/model"
	"gokomodo-assesment/app/domain/repository"
)

type SellerUsecase struct {
	repoJWT     repository.JWTRepository
	repoSeller  repository.SellerRepository
	repoProduct repository.ProductRepository
	repoOrder   repository.OrderRepository
}

func NewSellerUsecase(rj repository.JWTRepository, rs repository.SellerRepository, rp repository.ProductRepository, ro repository.OrderRepository) SellerPort {
	return &SellerUsecase{
		repoJWT:     rj,
		repoSeller:  rs,
		repoProduct: rp,
		repoOrder:   ro,
	}
}

func (u *SellerUsecase) Login(ctx context.Context, request *model.LoginRequestModel) (*model.LoginResponseModel, error) {
	// Get seller from env
	getSeller := &entity.Seller{
		Email:    request.Email,
		Password: "",
	}

	seller, err := u.repoSeller.GetSeller(getSeller)
	if err != nil {
		return nil, errors.New("invalid Email or Password")
	}

	// check if password valid
	err = u.repoJWT.CheckPassword(seller.Password, request.Password)
	if err != nil {
		return nil, errors.New("invalid Email or Password")
	}

	// generate jwt
	token, err := u.repoJWT.GenerateJWT(request.Email, "seller")
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

func (u *SellerUsecase) GetListProduct(ctx context.Context, sellerID string) ([]*model.ProductModel, error) {

	// Get product by seller id
	var response []*model.ProductModel

	listProduct, err := u.repoProduct.GetListProductBySellerID(sellerID)
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

func (u *SellerUsecase) AddNewProduct(ctx context.Context, product *model.ProductModel) error {

	// Set product entity
	newProduct := &entity.Product{
		SellerID:    product.SellerID,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
	}

	err := u.repoProduct.AddProduct(newProduct)
	if err != nil {
		return errors.New("failed to add product")
	}

	return nil
}

func (u *SellerUsecase) GetListOrder(ctx context.Context, sellerID string) ([]*model.SellerOrderModel, error) {

	var response []*model.SellerOrderModel
	listOrder, err := u.repoOrder.GetOrderBySellerID(sellerID)
	if err != nil {
		return nil, errors.New("failed to get list order")
	}

	for _, item := range listOrder {
		response = append(response, &model.SellerOrderModel{
			ID:                         item.ID,
			BuyerID:                    item.BuyerID,
			BuyerName:                  item.Buyer.Name,
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

func (u *SellerUsecase) AcceptOrder(ctx context.Context, request *model.AcceptOrderModel) error {

	order := &entity.Order{
		BaseModel: entity.BaseModel{
			ID: request.OrderID,
		},
		SellerID: request.SellerID,
	}

	err := u.repoOrder.AcceptOrder(order)
	if err != nil {
		return errors.New("failed to accept order")
	}

	return nil
}
