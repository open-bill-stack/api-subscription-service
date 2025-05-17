package subscription

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Service struct {
	repo       Repository
	grpcClient GrpcClient
}

func NewService(r Repository, g GrpcClient) *Service {
	return &Service{
		repo:       r,
		grpcClient: g,
	}
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Subscription, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, item *Subscription) (*Subscription, error) {
	statusExistUserID, errExistUserID := s.grpcClient.UserExistsByID(ctx, item.UserID.Bytes)
	if errExistUserID != nil {
		return nil, errExistUserID
	}
	if !statusExistUserID {
		return nil, fmt.Errorf("userID not exist")
	}
	statusExistTariffID, errExistTariffID := s.grpcClient.TariffExistsByID(ctx, item.TariffID.Bytes)
	if errExistTariffID != nil {
		return nil, errExistTariffID
	}
	if !statusExistTariffID {
		return nil, fmt.Errorf("tariff not exist")
	}
	return s.repo.Create(ctx, item)
}
func (s *Service) UpdateByID(ctx context.Context, item *Subscription) (*Subscription, error) {
	return s.repo.UpdateByID(ctx, item)
}
func (s *Service) DeleteByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.repo.DeleteByID(ctx, id)
}
func (s *Service) List(ctx context.Context) ([]Subscription, error) {
	return s.repo.List(ctx)
}
func (s *Service) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.repo.ExistsByID(ctx, id)
}

func (s *Service) DeleteByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	return s.repo.DeleteByUserID(ctx, userID)
}
func (s *Service) DeleteByTariffID(ctx context.Context, tariffID uuid.UUID) (bool, error) {
	return s.repo.DeleteByTariffID(ctx, tariffID)
}
