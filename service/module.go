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
	fx.Provide(func(repo repocontract.IProductRepository, nat repocontract.INatsRepository,
		dbRepo repocontract.IDBRepository, redisRepo repocontract.IRedisRepository) servicecontract.IProductService {
		return NewProductService(repo, nat, dbRepo, redisRepo)
	}),
)
