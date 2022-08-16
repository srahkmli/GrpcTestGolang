package cmd

import (
	"context"
	"log"
	"micro/client/elk"
	"micro/client/postgres"
	"micro/config"
	"micro/pkg/logger"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var migrateCMD = cobra.Command{
	Use:     "migrate",
	Long:    "migrate database strucutures. This will migrate tables",
	Aliases: []string{"m"},
	Run:     Runner.Migrate,
}

// migrate database with fake data
func (c *command) Migrate(cmd *cobra.Command, args []string) {
	fx.New(
		fx.Provide(elk.NewLogStash),
		fx.Provide(postgres.NewPostgres),
		fx.Invoke(config.InitConfigs),
		fx.Invoke(logger.InitGlobalLogger),
		fx.Invoke(migrateReducer(args)),
	).Start(context.TODO())
}

func migrateReducer(args []string) func(fx.Lifecycle, *pg.DB) {
	return func(l fx.Lifecycle, d *pg.DB) {
		migrate(l, d, args)
	}
}

func migrate(lc fx.Lifecycle, db *pg.DB, args []string) {
	lc.Append(fx.Hook{OnStart: func(c context.Context) error {
		oldVersion, newVersion, err := migrations.Run(db, args...)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		if newVersion != oldVersion {
			log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
		} else {
			log.Printf("version is %d\n", oldVersion)
		}
		return nil
	}})
}
