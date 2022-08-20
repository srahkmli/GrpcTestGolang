package repository

import (
	"context"
	"github.com/go-pg/pg/v10"
	"go.uber.org/fx"
	"log"
	"micro/client/broker"
	"micro/client/jtrace"
	"micro/client/redis"
	"micro/model"
	repocontract "micro/repository_contract"
)

type DBRepository struct {
	nats  broker.NatsBroker
	redis redis.Store
	db    *pg.DB
}

type DBRepositoryParams struct {
	fx.In

	Nats  broker.NatsBroker
	Redis redis.Store
	DB    *pg.DB
}

func NewDBRepository(params DBRepositoryParams) repocontract.IDBRepository {
	return &DBRepository{
		nats:  params.Nats,
		redis: params.Redis,
		db:    params.DB,
	}
}
func (b *DBRepository) InsertModel(ctx context.Context, m model.ProductModel) error {
	span, ctx := jtrace.T().SpanFromContext(ctx, "DBRepo[InsertModel]")
	defer span.Finish()
	_, err := b.db.Model(&m).Insert()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (b *DBRepository) SelectModel(ctx context.Context, point model.PointModel) (model.ProductModel, error) {
	span, ctx := jtrace.T().SpanFromContext(ctx, "DBRepo[SelectModel]")
	defer span.Finish()

	var res model.ProductModel
	if err := b.db.Model(&res).Where("name = ?", point.Point).Select(); err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}
