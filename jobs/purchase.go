package jobs

import (
	"context"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"micro/client/broker"
)

func StoreProductListener(lc fx.Lifecycle, Broker broker.NatsBroker) {
	subject := "StoreProduct"
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		Broker.Subscribe(subject, func(resp *nats.Msg) {
			zap.L().Info("I received a message")
			resp.Respond([]byte("message received"))
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
