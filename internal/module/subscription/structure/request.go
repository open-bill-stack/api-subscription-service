package structure

type CreateSubscription struct {
	UserID   string `json:"user_id" validate:"required"`
	TariffID string `json:"tariff_id" validate:"required"`
}
