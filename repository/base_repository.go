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

type BaseRepository struct {
	nats  broker.NatsBroker
	redis redis.Store
	db    *pg.DB
}

type BaseRepositoryParams struct {
	fx.In

	// Nats  broker.NatsBroker
	// Redis redis.Store
	// DB    *pg.DB
}

func NewBaseRepository(params BaseRepositoryParams) repocontract.IBaseRepository {
	return &BaseRepository{
		// nats:  params.Nats,
		// redis: params.Redis,
		// db:    params.DB,
	}
}

func (b *BaseRepository) StoreBaseModel(m model.BaseModel1) error {
	res, err := b.db.Model(&m).Insert()
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println(res)
	}
	return b.redis.Set(context.TODO(), config.C().Service.Name+":"+m.UserID, strconv.Itoa(int(m.Code)), time.Second*90)
}

func (b *BaseRepository) NotifySomeone(m model.BaseModel2) error {
	return b.nats.Publish(context.TODO(), "sample subject", m.Data)
}
