package service

import (
	"github.com/sirupsen/logrus"
	"github.com/srahkmli/grpcTest/model"
	"github.com/srahkmli/grpcTest/repository_contract"
)

type ProductImpl struct {
	productRepo repository.ProductRepository
}

func NewProduct(productRepo repository.ProductRepository) *ProductImpl {
	productService := ProductImpl{
		productRepo: productRepo,
	}
	return &productService
}
func (s *ProductImpl) AddProduct(product model.Product) error {
	logrus.Infof("add a new product %+v\n", product)
	err := s.productRepo.UpsertProduct(&product)
	if err != nil {
		logrus.Errorf("Failed to add a new product due to %v", err)
		return err
	}
	return nil
}
func (s *ProductImpl) FindProduct(key string) *model.Product {
	p := s.productRepo.FindProduct(key)
	return &p
}
