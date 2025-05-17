package subscription

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Subscription struct {
	ID       pgtype.UUID `db:"id"`      // UUID
	UserID   pgtype.UUID `db:"user_id"` // uuid користувача
	TariffID pgtype.UUID `db:"tariff_id"`
}
