package kinkajou

import (
	"api-subscription-service/internal/service/config"
	"api-subscription-service/internal/service/kinkajou/router"
	"context"
	"github.com/open-bill-stack/kinkajou"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	EventApp    *kinkajou.App
	AQMPChannel *amqp.Channel
	Log         *zap.Logger
	Config      *config.Config
	Router      []router.Router `group:"eventRoutes"`
}
type Result struct {
	fx.Out
	EventApp *kinkajou.App
}

func NewEventApp() (Result, error) {
	return Result{
		EventApp: kinkajou.New(),
	}, nil
}

func RunEventApp(lc fx.Lifecycle, p Params) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, r := range p.Router {
				r.Register(p.EventApp)
			}
			// Create the queue
			if _, err := p.AQMPChannel.QueueDeclare(
				"subscription.service", // name
				false,                  // durable
				false,                  // delete when unused
				false,                  // exclusive
				false,                  // no-wait
				nil,                    // arguments
			); err != nil {
				return err
			}
			// Bind the queue to the exchange
			if err := p.AQMPChannel.QueueBind(
				"subscription.service",
				"user.*",      // routing key
				"user.events", // name
				false,
				nil,
			); err != nil {
				return err
			}
			if err := p.AQMPChannel.QueueBind(
				"subscription.service",
				"tariff.*",      // routing key
				"tariff.events", // name
				false,
				nil,
			); err != nil {
				return err
			}
			// Subscribe to the queue
			ch, err := p.AQMPChannel.Consume(
				"subscription.service",
				"",
				false, // autoAck
				false, // exclusive
				false, // noLocal
				false, // noWait
				nil,   // args
			)
			if err != nil {
				return err
			}
			go func() {
				if err := p.EventApp.Listen(ch); err != nil {
					p.Log.Panic("Error starting Fiber server:", zap.Error(err))
				}
			}()
			return nil
		},
	})

}

var Module = fx.Module(
	"EventAppModule",
	fx.Provide(
		NewEventApp,
	),
	fx.Invoke(
		RunEventApp,
	),
)
