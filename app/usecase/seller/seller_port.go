package seller

import (
	"context"
	"gokomodo-assesment/app/domain/model"
)

type SellerPort interface {
	Login(ctx context.Context, request *model.LoginRequestModel) (*model.LoginResponseModel, error)
	GetListProduct(ctx context.Context, sellerID string) ([]*model.ProductModel, error)
	AddNewProduct(ctx context.Context, product *model.ProductModel) error
	GetListOrder(ctx context.Context, sellerID string) ([]*model.SellerOrderModel, error)
	AcceptOrder(ctx context.Context, request *model.AcceptOrderModel) error
}
