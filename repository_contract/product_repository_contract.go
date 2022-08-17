package repocontract

import (
	"context"
	"micro/model"
)

type IProductRepository interface {
	StoreProductModel(context.Context, model.ProductModel) error
	NotifyPurchase(context.Context, model.ProductModel) (model.ProductModel, error)
	GetProductModel(context.Context, model.PointModel, string) (model.ProductModel, string, error)
}
