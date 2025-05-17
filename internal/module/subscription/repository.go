package subscription

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, item *Subscription) (*Subscription, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Subscription, error)
	UpdateByID(ctx context.Context, item *Subscription) (*Subscription, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, error)
	List(ctx context.Context) ([]Subscription, error)
	ExistsByID(ctx context.Context, id uuid.UUID) (bool, error)

	DeleteByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
	DeleteByTariffID(ctx context.Context, tariffID uuid.UUID) (bool, error)
}
