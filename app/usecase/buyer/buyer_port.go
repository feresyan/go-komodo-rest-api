package buyer

import (
	"context"
	"gokomodo-assesment/app/domain/model"
)

type BuyerPort interface {
	Login(ctx context.Context, request *model.LoginRequestModel) (*model.LoginResponseModel, error)
	GetListProduct(ctx context.Context) ([]*model.ProductModel, error)
	AddOrder(ctx context.Context, request *model.AddOrderModel) error
	GetListOrder(ctx context.Context, buyerID string) ([]*model.BuyerOrderModel, error)
}
