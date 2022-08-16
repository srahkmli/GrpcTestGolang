package controller

import (
	controller "micro/controller/base"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(controller.NewBaseController),
)
