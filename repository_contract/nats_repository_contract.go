package repocontract

import (
	"context"
	"micro/model"
)

type INatsRepository interface {
	StoreProductModel(context.Context, model.ProductModel) error
	ReturnPurchaseModel(context.Context, model.PointModel) (model.ProductModel, error)
}
