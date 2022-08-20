package repository

import (
	"context"
	"github.com/go-pg/pg/v10"
	"go.uber.org/fx"
	"log"
	"micro/client/broker"
	"micro/client/jtrace"
	"micro/client/redis"
	"micro/config"
	"micro/model"
	repocontract "micro/repository_contract"
	"time"
)

type RedisRepository struct {
	nats  broker.NatsBroker
	redis redis.Store
	db    *pg.DB
}

type RedisRepositoryParams struct {
	fx.In

	Nats  broker.NatsBroker
	Redis redis.Store
	DB    *pg.DB
}

func NewRedisRepository(params RedisRepositoryParams) repocontract.IRedisRepository {
	return &RedisRepository{
		nats:  params.Nats,
		redis: params.Redis,
		db:    params.DB,
	}
}
func (b *RedisRepository) SetCache(ctx context.Context, productModel model.ProductModel) error {
	span, ctx := jtrace.T().SpanFromContext(ctx, "RedisRepo[SetCache]")
	defer span.Finish()

	var res model.ProductModel
	cacheKey := config.C().Service.Name + ":" + productModel.Name
	if err := b.redis.Set(ctx, cacheKey, &res, time.Minute*90); err != nil {
		log.Println(err)
	}
	return nil

}
func (b *RedisRepository) GetCache(ctx context.Context, p model.PointModel) (model.ProductModel, error) {
	span, ctx := jtrace.T().SpanFromContext(ctx, "RedisRepo[GetCache]")
	defer span.Finish()

	var res model.ProductModel
	cacheKey := config.C().Service.Name + ":" + p.Point
	err := b.redis.Get(ctx, cacheKey, &res)
	if err == nil {
		log.Printf("from redis name is : %s  , qty is %v", res.Name, res.Qty)
		return res, err
	}
	return res, err
}
