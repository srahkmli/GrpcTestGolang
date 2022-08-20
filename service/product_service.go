package service

import (
	"context"
	"micro/client/jtrace"
	"micro/model"
	repocontract "micro/repository_contract"
	"regexp"
)

type ProductService struct {
	productRepo repocontract.IProductRepository
	natsRepo    repocontract.INatsRepository
	dbRepo      repocontract.IDBRepository
	redisRepo   repocontract.IRedisRepository
}

func (p ProductService) Validate(name string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	return isAlpha(name)
}

func NewProductService(repo repocontract.IProductRepository, nat repocontract.INatsRepository,
	dbRepo repocontract.IDBRepository, redisRepo repocontract.IRedisRepository) ProductService {
	return ProductService{
		productRepo: repo,
		natsRepo:    nat,
		dbRepo:      dbRepo,
		redisRepo:   redisRepo,
	}
}

func (p ProductService) SetProcess(ctx context.Context, product model.ProductModel) error {
	span, ctx := jtrace.T().SpanFromContext(ctx, "service[SetProcess]")
	defer span.Finish()

	err := p.natsRepo.StoreProductModel(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (p ProductService) GetProcess(ctx context.Context, point model.PointModel) (model.ProductModel, error) {
	span, ctx := jtrace.T().SpanFromContext(ctx, "service[GetProcess]")
	defer span.Finish()

	result, err := p.redisRepo.GetCache(ctx, point)
	if err == nil {
		return result, err
	}
	result, err = p.dbRepo.SelectModel(ctx, point)
	if err != nil {
		return result, err
	}
	p.redisRepo.SetCache(ctx, result)

	return result, nil
}
