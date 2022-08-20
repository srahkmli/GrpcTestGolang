package repocontract

import (
	"context"
	"micro/model"
)

type IDBRepository interface {
	InsertModel(context.Context, model.ProductModel) error
	SelectModel(context.Context, model.PointModel) (model.ProductModel, error)
}
