package service_contract

import "github.com/srahkmli/grpcTest/model"

type ProductService interface {
	AddProduct(model.Product) error
	FindProduct(string) *model.Product
	RemoveProduct(int32) error
}
