package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewBaseRepository),
	fx.Provide(NewProductRepository),
	fx.Provide(NewNatsRepository),
	fx.Provide(NewDBRepository),
	fx.Provide(NewRedisRepository),
)
