package app

import (
	"context"
	"fmt"
	"log"
	"micro/app/server"
	"micro/client/broker"
	"micro/client/elk"
	"micro/client/jtrace"
	"micro/client/postgres"
	"micro/client/redis"
	"micro/jobs"
	"micro/pkg/logger"
	"os"
	"time"

	//	"micro/client/broker"
	//	"micro/client/jtrace"
	//	"micro/client/postgres"
	//	"micro/client/redis"
	"micro/config"
	"micro/controller"

	//	"micro/pkg/logger"
	"micro/repository"
	"micro/service"

	"go.uber.org/fx"
)

// StartApplication func
func Start() {
	fmt.Println("\n\n--------------------------------")
	// if go code crashed we get error and line
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// init configs

	for {
		fxNew := fx.New(
			fx.Provide(broker.NewNats),
			fx.Provide(redis.NewRedis),
			fx.Provide(postgres.NewPostgres),
			fx.Provide(elk.NewLogStash),
			service.Module,
			repository.Module,
			controller.Module,
			fx.Provide(server.New),
			fx.Invoke(config.InitConfigs),
			fx.Invoke(logger.InitGlobalLogger),
			fx.Invoke(jtrace.InitGlobalTracer),
			fx.Invoke(jobs.InitProductListener),
			fx.Invoke(serve),
		)
		startCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := fxNew.Start(startCtx); err != nil {
			log.Println(err)
			break
		}
		if val := <-fxNew.Done(); val == os.Interrupt {
			break
		}

		stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := fxNew.Stop(stopCtx); err != nil {
			log.Println(err)
			break
		}
	}
}

func serve(lc fx.Lifecycle, server server.IServer) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return server.ListenAndServe()
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})
}
