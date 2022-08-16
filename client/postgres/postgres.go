package postgres

import (
	"context"
	"log"
	"micro/config"

	"github.com/go-pg/pg/v10"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Connect method job is connect to postgres database and check migration
func NewPostgres(lc fx.Lifecycle) *pg.DB {
	var err error
	db := pg.DB{}
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			db = *pg.Connect(&pg.Options{
				User:                  config.C().Postgres.Username,
				Password:              config.C().Postgres.Password,
				Addr:                  config.C().Postgres.Host,
				Database:              config.C().Postgres.Schema,
				RetryStatementTimeout: true,
			})
			if err = db.Ping(context.Background()); err != nil {
				zap.L().Error(err.Error())
				return err
			}
			log.Printf("postgres database loaded successfully \n")
			return nil
		},
		OnStop: func(c context.Context) error {
			log.Printf("postgres database connection closed \n")
			return db.Close()
		},
	})

	return &db
}
