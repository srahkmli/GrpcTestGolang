package repocontract

import (
	"context"
	"micro/model"
)

type IRedisRepository interface {
	GetCache(context.Context, model.PointModel) (model.ProductModel, error)
	SetCache(context.Context, model.ProductModel) error
}
