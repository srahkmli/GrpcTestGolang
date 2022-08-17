package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"micro/client/jtrace"
	"micro/model"
	repocontract "micro/repository_contract"
	"regexp"
)

type ProductService struct {
	productRepo repocontract.IProductRepository
	cacheKey    string
}

func (p ProductService) Validate(name string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	return isAlpha(name)
}

func NewProductService(repo repocontract.IProductRepository) ProductService {
	return ProductService{
		productRepo: repo,
		cacheKey:    "",
	}
}

func (p ProductService) SetProcess(ctx context.Context, m model.ProductModel) (model.PurchaseModel, error) {
	span, ctx := jtrace.T().SpanFromContext(ctx, "srvis[SetProcess]")
	defer span.Finish()

	var result model.PurchaseModel
	err := p.productRepo.StoreProductModel(ctx, m)
	if err != nil {
		return result, err
	}

	zap.L().Debug("stored ok")
	resProduct, err := p.productRepo.NotifyPurchase(ctx, m)
	if err != nil {
		return result, err
	}

	result.Data = resProduct.Name
	return result, nil
}

func (p ProductService) GetProcess(ctx context.Context, m model.PointModel) (model.PurchaseModel, error) {
	span, ctx := jtrace.T().SpanFromContext(ctx, "service[GetProcess]")
	defer span.Finish()
	log.Printf("current chachKey is %v ", p.cacheKey)
	res, cacheKey, err := p.productRepo.GetProductModel(ctx, m, p.cacheKey)
	if err != nil {
		return model.PurchaseModel{}, err
	}
	p.cacheKey = cacheKey

	log.Printf("update chachKey is %v ", p.cacheKey)
	result := model.PurchaseModel{Data: fmt.Sprintf("%s - %d", res.Name, res.Name)}
	return result, nil
}
