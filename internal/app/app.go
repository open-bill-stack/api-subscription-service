package app

import (
	"api-subscription-service/internal/module/subscription"
	"api-subscription-service/internal/service/config"
	"api-subscription-service/internal/service/database"
	"api-subscription-service/internal/service/fiber"
	"api-subscription-service/internal/service/fiber/middleware"
	"api-subscription-service/internal/service/grpc"
	"api-subscription-service/internal/service/jwt"
	"api-subscription-service/internal/service/kinkajou"
	"api-subscription-service/internal/service/logger"
	"api-subscription-service/internal/service/rabbitmq"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Run(cmd *cobra.Command) {
	fx.New(
		fx.Provide(func() *cobra.Command { return cmd }),

		logger.Module,
		config.Module,
		database.Module,
		jwt.Module,

		// global middleware
		middleware.Module,

		// module
		subscription.Module,

		rabbitmq.Module,
		kinkajou.Module,
		// fiber
		fiber.Module,

		// grpc
		grpc.ModuleClient,
	).Run()
}
