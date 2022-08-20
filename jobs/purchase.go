package jobs

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"micro/client/broker"
	"micro/model"
	servicecontract "micro/service_contract"
)

type StoreProductListenerParams struct {
	fx.In
	ProductService servicecontract.IProductService
}

func InitProductListener(lc fx.Lifecycle, Broker broker.NatsBroker, param StoreProductListenerParams) {
	param.StoreProductListener(lc, Broker)
}
func (p *StoreProductListenerParams) StoreProductListener(lc fx.Lifecycle, Broker broker.NatsBroker) {
	subject := "StoreProduct"
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		Broker.Subscribe(subject, func(product *model.ProductModel) {
			zap.L().Info("I received a message")
			p.ProductService.SetProcess(ctx, *product)
		})
		return nil
	}})
}
func PurchaseListener(Broker broker.NatsBroker) {
	//subject := "ReturnPurchase"
	//Broker.Subscribe(subject, func(m *nats.Msg) {
	//	fmt.Printf("Received a message: %s\n", string(m.Data))
	//})
}
