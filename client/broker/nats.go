package broker

import (
	"context"
	"log"
	"micro/config"
	"time"

	nats "github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type nts struct {
	nc *nats.EncodedConn
}

func NewNats(lc fx.Lifecycle) NatsBroker {
	nts := nts{}
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			if err := nts.connect(*config.C()); err != nil {
				log.Printf("nats message broker failed to start \n")
				return err
			}
			log.Printf("nats message broker loaded successfully \n")
			return nil
		},
		OnStop: func(c context.Context) error {
			nts.Encoded().Close()
			log.Printf("nats message broker shutdown \n")
			return nil
		},
	})
	return &nts
}

// Connect nats broker
func (n *nts) connect(conf config.Config) error {
	var err error

	var conn *nats.Conn
	opts := nats.Options{
		Name:           conf.Service.Name,
		Secure:         conf.Nats.Auth,
		User:           conf.Nats.Username,
		Password:       conf.Nats.Password,
		Servers:        conf.Nats.Endpoints,
		PingInterval:   time.Second * 60,
		AllowReconnect: conf.Nats.AllowReconnect,
		MaxReconnect:   conf.Nats.MaxReconnect,
		ReconnectWait:  time.Duration(conf.Nats.ReconnectWait) * time.Second,
		Timeout:        time.Duration(conf.Nats.Timeout) * time.Second,
		AsyncErrorCB:   n.ErrorReporter(zap.L()),
	}

	// try to connect to nats message broker
	conn, err = opts.Connect()
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	n.nc, err = nats.NewEncodedConn(conn, conf.Nats.Encoder)
	return err
}

// Encoded get Connection
func (n *nts) Encoded() *nats.EncodedConn {
	return n.nc
}

// Publish new message
func (n *nts) Publish(ctx context.Context, subject string, value interface{}) error {
	if err := n.Encoded().Publish(subject, &value); err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}

// SendChan, send a untyped chan
func (n *nts) SendChan(subject string, ch chan interface{}) error {
	return n.Encoded().BindSendChan(subject, ch)
}

// SendByContext, send a request and get response with spesific context
func (n *nts) SendByContext(ctx context.Context, subject string, req interface{}, resp interface{}) error {
	if err := n.RequestWithContext(ctx, subject, req, resp); err != nil {
		zap.L().Info("msg", zap.String("err", err.Error()))
		return err
	}
	return nil
}

// RequestWithReply, send a request and get response of request
// Then call Flush
func (n *nts) RequestWithReply(subject string, req interface{}, resp string) error {
	if err := n.Encoded().PublishRequest(subject, resp, req); err != nil {
		zap.L().Info("msg", zap.String("err", err.Error()))
		return err
	}

	if err := n.Encoded().Flush(); err != nil {
		zap.L().Info("msg", zap.String("err", err.Error()))
		return err
	}

	return nil
}

// Subscribe, start to subscribe to a subject
func (n *nts) Subscribe(subject string, callBack func(resp *nats.Msg)) (*nats.Subscription, error) {
	sub, err := n.Encoded().Subscribe(subject, callBack)
	if err != nil {
		zap.L().Info("msg", zap.String("err", err.Error()))
		return nil, err
	}

	return sub, nil
}

// RecvChan, BindRecvChan
func (n *nts) RecvChan(subject string, ch chan interface{}) (*nats.Subscription, error) {
	sub, err := n.Encoded().BindRecvChan(subject, ch)

	if err != nil {
		zap.L().Info("msg", zap.String("err", err.Error()))
		return nil, err
	}

	return sub, nil
}

// RecvGroup, connect to subject in group mode
func (n *nts) RecvGroup(subject, queue string, callBack nats.Handler) (*nats.Subscription, error) {
	sub, err := n.Encoded().QueueSubscribe(subject, queue, callBack)
	if err != nil {
		zap.L().Info("msg", zap.String("err", err.Error()))
		return nil, err
	}

	return sub, nil
}

// ErrorReporter, when nats has error
func (n *nts) ErrorReporter(log *zap.Logger) nats.ErrHandler {
	return func(_ *nats.Conn, sub *nats.Subscription, err error) {
		pendingMsgs, pendingBytes, _ := sub.Pending()
		droppedMsgs, _ := sub.Dropped()
		maxMsgs, maxBytes, _ := sub.PendingLimits()

		log.Error(err.Error(),
			zap.String("subject", sub.Subject),
			zap.String("queue", sub.Queue),
			zap.Int("pending_msgs", pendingMsgs),
			zap.Int("pending_bytes", pendingBytes),
			zap.Int("max_msgs_pending", maxMsgs),
			zap.Int("max_bytes_pending", maxBytes),
			zap.Int("dropped_msgs", droppedMsgs),
			zap.String("message", "Error while consuming from nats"),
		)
	}
}

func (n *nts) RequestWithContext(ctx context.Context, subject string, v interface{}, vPtr interface{}) error {
	err := n.Encoded().RequestWithContext(ctx, subject, v, vPtr)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	return nil
}
