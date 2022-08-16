package controller

import (
	"go.uber.org/fx"
	controller "micro/controller/base"
	productController "micro/controller/product"
)

var Module = fx.Options(
	fx.Provide(controller.NewBaseController),
	fx.Provide(productController.NewProductController),
)
