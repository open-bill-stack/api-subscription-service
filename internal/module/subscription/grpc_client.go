package subscription

import (
	"context"
	"github.com/google/uuid"
)

type GrpcClient interface {
	UserExistsByID(ctx context.Context, userID uuid.UUID) (bool, error)
	TariffExistsByID(ctx context.Context, userID uuid.UUID) (bool, error)
}
