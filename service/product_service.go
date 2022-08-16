package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	redicontract "micro/client/redis"
	"micro/model"
	repocontract "micro/repository_contract"
	"regexp"
	"time"
)

type ProductService struct {
	productRepo repocontract.IProductRepository
	redis       redicontract.Store
}

func (p ProductService) Validate(name string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	return isAlpha(name)
}

func NewProductService(repo repocontract.IProductRepository, store redicontract.Store) ProductService {
	return ProductService{
		productRepo: repo,
		redis:       store,
	}
}

func (p ProductService) Process(ctx context.Context, m model.ProductModel) (model.PurchaseModel, error) {
	var result model.PurchaseModel
	var cacheKey = fmt.Sprintf("%s - %d", m.Name, m.Qty)

	if err := p.redis.Get(context.Background(), cacheKey, result); err == nil {
		return result, nil
	}
	zap.L().Debug("redis get ok")
	if err := p.productRepo.StoreProductModel(m); err != nil {
		return result, err
	}
	zap.L().Debug("stored ok")
	if err := p.redis.Set(context.Background(), cacheKey, result, time.Duration(3*time.Now().Day())); err != nil {
		zap.L().Error(err.Error())
		return result, nil
	}
	zap.L().Debug("redis set ok")
	result.Data = fmt.Sprintf("you purchased a  %s - %d", m.Name, m.Qty)

	if err := p.productRepo.NotifyPurchase(result); err != nil {
		return result, err
	}
	return result, nil
}
