package repository

import (
	"context"
	"log"
	"micro/client/broker"
	"micro/client/redis"
	"micro/config"
	"micro/model"
	repocontract "micro/repository_contract"
	"strconv"
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

func (b *ProductRepository) StoreProductModel(m model.ProductModel) error {
	res, err := b.db.Model(&m).Insert()
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println(res)
	}
	return b.redis.Set(context.TODO(), config.C().Service.Name+":"+m.Name, strconv.Itoa(int(m.Qty)), time.Second*90)
}

func (b *ProductRepository) NotifyPurchase(m model.PurchaseModel) error {
	return b.nats.Publish(context.TODO(), "sample subject", m.Data)
}
