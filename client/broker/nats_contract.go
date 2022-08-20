package broker

import (
	"context"
	"micro/model"

	nats "github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// NatsBroker interface
type NatsBroker interface {
	Encoded() *nats.EncodedConn
	Publish(ctx context.Context, subject string, value interface{}) error
	SendChan(subject string, ch chan interface{}) error
	SendByContext(ctx context.Context, subject string, req interface{}, resp interface{}) error
	RequestWithReply(subject string, req interface{}, resp string) error
	Subscribe(subject string, callBack func(product *model.ProductModel)) (*nats.Subscription, error)
	RecvChan(subject string, ch chan interface{}) (*nats.Subscription, error)
	RecvGroup(subject, queue string, callBack nats.Handler) (*nats.Subscription, error)
	ErrorReporter(log *zap.Logger) nats.ErrHandler
	RequestWithContext(ctx context.Context, subject string, v interface{}, vPtr interface{}) error
}
