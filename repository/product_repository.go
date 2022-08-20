package repository

import (
	"context"
	"fmt"
	"log"
	"micro/client/broker"
	"micro/client/jtrace"
	"micro/client/redis"
	"micro/config"
	"micro/model"
	repocontract "micro/repository_contract"
	"time"

	"github.com/go-pg/pg/v10"
	"go.uber.org/fx"
)

type ProductRepository struct {
	nats  broker.NatsBroker
	redis redis.Store
	db    *pg.DB
}

type ProductRepositoryParams struct {
	fx.In

	Nats  broker.NatsBroker
	Redis redis.Store
	DB    *pg.DB
}

func NewProductRepository(params ProductRepositoryParams) repocontract.IProductRepository {
	return &ProductRepository{
		nats:  params.Nats,
		redis: params.Redis,
		db:    params.DB,
	}
}

func (b *ProductRepository) StoreProductModel(ctx context.Context, m model.ProductModel) error {
	span, ctx := jtrace.T().SpanFromContext(ctx, "repo[StoreProductModel]")
	defer span.Finish()
	_, err := b.db.Model(&m).Insert()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("stored successfully")
	return nil
}

func (b *ProductRepository) NotifyPurchase(ctx context.Context, m model.ProductModel) (model.PurchaseModel, error) {
	span, ctx := jtrace.T().SpanFromContext(ctx, "repo[NotifyPurchase]")
	defer span.Finish()

	result := model.PurchaseModel{Data: fmt.Sprintf("Hello %s - %d", m.Name, m.Qty)}

	return result, b.nats.Publish(ctx, "product subject", result)
}

func (b *ProductRepository) GetProductModel(ctx context.Context, m model.PointModel) (model.ProductModel, error) {
	span, ctx := jtrace.T().SpanFromContext(ctx, "repo[GetProductModel]")
	defer span.Finish()

	var res model.ProductModel

	cacheKey := config.C().Service.Name + ":" + m.Point
	err := b.redis.Get(ctx, cacheKey, &res)
	if err == nil {
		log.Printf("from redis name is : %s  , qty is %v", res.Name, res.Qty)
		return res, b.nats.Publish(context.TODO(), "product subject", res)
	}
	log.Printf("redis err : %v", err)

	log.Println("DB Called")
	if err := b.db.Model(&res).Where("name = ?", m.Point).Select(); err != nil {
		log.Println(err)
		return res, err
	}

	log.Printf("setting redis")

	if err := b.redis.Set(ctx, cacheKey, &res, time.Minute*90); err != nil {
		log.Println(err)
	}

	return res, b.nats.Publish(context.TODO(), "product subject", res)
}
