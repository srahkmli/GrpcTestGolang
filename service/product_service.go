package service

import (
	"fmt"
	"google.golang.org/genproto/googleapis/cloud/redis/v1"
	"micro/model"
	repocontract "micro/repository_contract"
	"regexp"
)

type ProductService struct {
	productRepo repocontract.IProductRepository
	redis       redis.UnimplementedCloudRedisServer
}

func (p ProductService) Validate(name string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	return isAlpha(name)
}

func NewProductService(repo repocontract.IProductRepository) ProductService {
	return ProductService{
		productRepo: repo,
	}
}

func (p ProductService) Process(m model.ProductModel) (model.PurchaseModel, error) {
	result := model.PurchaseModel{Data: fmt.Sprintf("you purchased a  %s - %d", m.Name, m.Qty)}
	if err := p.productRepo.StoreProductModel(m); err != nil {
		return result, err
	}
	if err := p.productRepo.NotifyPurchase(result); err != nil {
		return result, err
	}
	return result, nil
}
