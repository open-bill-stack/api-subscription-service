package middleware

import (
	"api-subscription-service/internal/service/fiber/middleware/cors"
	"api-subscription-service/internal/service/fiber/middleware/healthcheck"
	"api-subscription-service/internal/service/fiber/middleware/jwt"
	middlewareRecover "api-subscription-service/internal/service/fiber/middleware/recover"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"MiddlewareModule",
	fx.Provide(
		healthcheck.NewMiddleware,
		middlewareRecover.NewMiddleware,
		cors.NewMiddleware,
		jwt.NewService,
	),
)

//
