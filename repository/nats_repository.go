package repository

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"go.uber.org/fx"
	"log"
	"micro/client/broker"
	"micro/client/jtrace"
	"micro/client/redis"
	"micro/model"
	repocontract "micro/repository_contract"
)

type NatsRepository struct {
	nats  broker.NatsBroker
	redis redis.Store
	db    *pg.DB
}

type NatsRepositoryParams struct {
	fx.In

	Nats  broker.NatsBroker
	Redis redis.Store
	DB    *pg.DB
}

func NewNatsRepository(params NatsRepositoryParams) repocontract.INatsRepository {
	return &NatsRepository{
		nats:  params.Nats,
		redis: params.Redis,
		db:    params.DB,
	}
}

func (n *NatsRepository) StoreProductModel(ctx context.Context, product model.ProductModel) error {
	span, _ := jtrace.T().SpanFromContext(ctx, "NatsRepo[StoreProduct]")
	defer span.Finish()

	subject := "StoreProduct"
	_, err := n.db.Model(&product).Insert()
	if err != nil {
		log.Println(err)
		return err
	}
	if err := n.nats.Publish(ctx, subject, product); err != nil {
		return fmt.Errorf("publishing on nats failed: %w", err)
	}
	return nil
}

func (n *NatsRepository) ReturnPurchaseModel(ctx context.Context, p model.PointModel) (model.ProductModel, error) {

	panic("not implemented")
	//	span, ctx := jtrace.T().SpanFromContext(ctx, "NatsRepo[ReturnPurchase]")
	//	defer span.Finish()
	//	subject := "ReturnPurchase"
	//
	//	var res model.ProductModel
	//	cacheKey := config.C().Service.Name + ":" + p.Point
	//	err := n.redis.Get(ctx, cacheKey, &res)
	//	if err == nil {
	//		log.Printf("from redis name is : %s  , qty is %v", res.Name, res.Qty)
	//		return res, err
	//	}
	//	log.Printf("redis err : %v", err)
	//
	//	log.Println("DB Called")
	//	if err := n.db.Model(&res).Where("name = ?", p.Point).Select(); err != nil {
	//		log.Println(err)
	//		return res, err
	//	}
	//	if err := n.redis.Set(ctx, cacheKey, &res, time.Minute*90); err != nil {
	//		log.Println(err)
	//	}
	//	if err := n.nats.Publish(ctx, subject, res); err != nil {
	//		return res, fmt.Errorf("publishing on nats failed: %w", err)
	//	}
	//	return res, nil
}
