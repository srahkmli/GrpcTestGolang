package servicecontract

import (
	"golang.org/x/net/context"
	"micro/model"
)

type IProductService interface {
	Validate(string) bool
	SetProcess(context.Context, model.ProductModel) (model.PurchaseModel, error)
	GetProcess(context.Context, model.PointModel) (model.PurchaseModel, error)
}
