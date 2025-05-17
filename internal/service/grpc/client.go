package grpc

import (
	"api-subscription-service/internal/service/config"
	pbTariff "api-subscription-service/internal/service/grpc/proto/tariff/v1"
	pbUser "api-subscription-service/internal/service/grpc/proto/user/v1"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ParamsClient struct {
	fx.In

	Log    *zap.Logger
	Config *config.Config
}

type ResultUserClient struct {
	fx.Out
	GrpcClient *pbUser.UserServiceClient
}
type ResultTariffClient struct {
	fx.Out
	GrpcClient *pbTariff.TariffServiceClient
}

func NewGrpcUserClientApp(p ParamsClient) (ResultUserClient, error) {
	conn, err := grpc.NewClient(
		p.Config.Service.UserGrpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("grpc client connection error", zap.Error(err))
		panic(err)
	}
	client := pbUser.NewUserServiceClient(conn)
	return ResultUserClient{
		GrpcClient: &client,
	}, nil
}
func NewGrpcTariffClientApp(p ParamsClient) (ResultTariffClient, error) {
	conn, err := grpc.NewClient(
		p.Config.Service.TariffGrpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("grpc client connection error", zap.Error(err))
		panic(err)
	}
	client := pbTariff.NewTariffServiceClient(conn)
	return ResultTariffClient{
		GrpcClient: &client,
	}, nil
}

var ModuleClient = fx.Module(
	"GrpcClientAppModule",
	fx.Provide(
		NewGrpcUserClientApp,
		NewGrpcTariffClientApp,
	),
)
