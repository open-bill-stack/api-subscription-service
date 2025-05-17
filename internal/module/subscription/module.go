package subscription

import "go.uber.org/fx"

var Module = fx.Module("subscription",
	fx.Provide(NewHttpHandler),
	fx.Provide(NewEventHandler),

	fx.Provide(NewService),
	fx.Provide(NewRepository),
	fx.Provide(NewGrpcClient),
)
