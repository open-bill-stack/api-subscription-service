package subscription

import (
	httpRouter "api-subscription-service/internal/service/fiber/router"
	grpcRouter "api-subscription-service/internal/service/grpc/router"
	eventRouter "api-subscription-service/internal/service/kinkajou/router"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Log     *zap.Logger
	Service *Service
}
type HttpResult struct {
	fx.Out

	Router httpRouter.Router `group:"httpRoutes"`
}
type GrpcResult struct {
	fx.Out

	Router grpcRouter.Router `group:"grpcRoutes"`
}

type EventResult struct {
	fx.Out

	Router eventRouter.Router `group:"eventRoutes"`
}
