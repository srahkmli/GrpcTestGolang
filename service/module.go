package service

import (
	repocontract "micro/repository_contract"
	servicecontract "micro/service_contract"

	redicontract "micro/client/redis"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(func(repo repocontract.IBaseRepository) servicecontract.IBaseService {
		return NewBaseService(repo)
	}),
	fx.Provide(func(repo repocontract.IProductRepository, store redicontract.Store) servicecontract.IProductService {
		return NewProductService(repo, store)
	}),
)
