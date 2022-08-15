package repository

import (
	"github.com/srahkmli/grpcTest/model"
)

type ProductRepository interface {
	UpsertProduct(*model.Product) error
	FindProduct(string) model.Product
	DeleteProduct(string) error
}
