package servicecontract

import (
	"golang.org/x/net/context"
	"micro/model"
)

type IProductService interface {
	Validate(string) bool
	Process(context.Context, model.ProductModel) (model.PurchaseModel, error)
}
