package subscription

import (
	pbTariff "api-subscription-service/internal/service/grpc/proto/tariff/v1"
	pbUser "api-subscription-service/internal/service/grpc/proto/user/v1"
	"context"
	"github.com/google/uuid"
)

type grpcClient struct {
	clientUser   pbUser.UserServiceClient
	clientTariff pbTariff.TariffServiceClient
}

func (g *grpcClient) UserExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	r, err := g.clientUser.ExistsByID(ctx, &pbUser.ExistsByIDRequest{UserId: id.String()})
	if err != nil {
		return false, err
	}
	return r.GetExists(), nil
}

func (g *grpcClient) TariffExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	r, err := g.clientTariff.ExistsByID(ctx, &pbTariff.ExistsByIDRequest{TariffId: id.String()})
	if err != nil {
		return false, err
	}
	return r.GetExists(), nil
}

func NewGrpcClient(clientUser *pbUser.UserServiceClient, clientTariff *pbTariff.TariffServiceClient) GrpcClient {
	return &grpcClient{
		clientUser:   *clientUser,
		clientTariff: *clientTariff,
	}
}
