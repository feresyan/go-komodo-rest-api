package model

type SellerOrderModel struct {
	ID                         string  `json:"id"`
	BuyerID                    string  `json:"buyer_id"`
	BuyerName                  string  `json:"buyer_name"`
	DeliveryDestinationAddress string  `json:"delivery_destination_address"`
	ProductID                  string  `json:"product_id"`
	ProductName                string  `json:"product_name"`
	Quantity                   int64   `json:"quantity"`
	Price                      float64 `json:"price"`
	TotalPrice                 float64 `json:"total_price"`
	Status                     string  `json:"status"`
}

type AcceptOrderModel struct {
	OrderID  string `json:"order_id"`
	SellerID string `json:"seller_id"`
}

type AddOrderModel struct {
	BuyerID                    string  `json:"buyer_id"`
	SellerID                   string  `json:"seller_id"`
	BuyerName                  string  `json:"buyer_name"`
	DeliverySourceAddress      string  `json:"delivery_source_address"`
	DeliveryDestinationAddress string  `json:"delivery_destination_address"`
	ProductID                  string  `json:"product_id"`
	Quantity                   int64   `json:"quantity"`
	Price                      float64 `json:"price"`
	TotalPrice                 float64 `json:"total_price"`
}

type BuyerOrderModel struct {
	ID                         string  `json:"id"`
	BuyerID                    string  `json:"buyer_id"`
	BuyerName                  string  `json:"buyer_name"`
	SellerID                   string  `json:"seller_id"`
	SellerName                 string  `json:"seller_name"`
	DeliverySourceAddress      string  `json:"delivery_source_address"`
	DeliveryDestinationAddress string  `json:"delivery_destination_address"`
	ProductID                  string  `json:"product_id"`
	ProductName                string  `json:"product_name"`
	Quantity                   int64   `json:"quantity"`
	Price                      float64 `json:"price"`
	TotalPrice                 float64 `json:"total_price"`
	Status                     string  `json:"status"`
}

type OrderResponseModel struct {
	ResponseCode string `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
}
