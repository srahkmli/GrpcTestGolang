package service

import (
	repocontract "micro/repository_contract"
	servicecontract "micro/service_contract"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(func(repo repocontract.IBaseRepository) servicecontract.IBaseService {
		return NewBaseService(repo)
	}),
	fx.Provide(func(repo repocontract.IProductRepository) servicecontract.IProductService {
		return NewProductService(repo)
	}),
)
