package controller

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/srahkmli/grpcTest/model"
	"github.com/srahkmli/grpcTest/pb"
	"github.com/srahkmli/grpcTest/service_contract"
)

type ProductController struct {
	service service_contract.ProductService
}

func (s *ProductController) SaveProduct(ctx context.Context, product *grpcTest.Product) (model.Product, error) {

	logrus.Printf("Recieved %s", product.GetName())
	newProduct := model.NewProduct(product.GetName(), product.GetQty())
	if err := s.service.AddProduct(newProduct); err != nil {
		logrus.Printf("Error accured due to : %v", err)
	}
	return newProduct, nil
}
func (s *ProductController) ListProduct(ctx context.Context, point *grpcTest.Point) *model.Product {
	logrus.Printf("Recieved %v", point.GetName())
	product := s.service.FindProduct(point.GetName())
	return product
}
