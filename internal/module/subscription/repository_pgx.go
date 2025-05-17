package subscription

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type sqlRepo struct {
	db *pgxpool.Pool
}

func (s sqlRepo) DeleteByUserID(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `DELETE FROM subscriptions WHERE user_id = $1`
	cmdTag, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return false, nil
	}
	return true, nil
}

func (s sqlRepo) DeleteByTariffID(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `DELETE FROM subscriptions WHERE tariff_id = $1`
	cmdTag, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete tariff: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return false, nil
	}
	return true, nil
}

func (s sqlRepo) Create(ctx context.Context, item *Subscription) (*Subscription, error) {
	query := `INSERT INTO subscriptions (user_id, tariff_id) VALUES ($1, $2) RETURNING id, user_id, tariff_id`
	row := s.db.QueryRow(ctx, query, item.UserID, item.TariffID)
	var result Subscription
	err := row.Scan(&result.ID, &result.UserID, &result.TariffID)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}
	return &result, nil
}

func (s sqlRepo) GetByID(ctx context.Context, id uuid.UUID) (*Subscription, error) {
	query := `SELECT * FROM subscriptions WHERE id = $1`
	var item Subscription
	err := s.db.QueryRow(ctx, query, id).Scan(&item.ID, &item.UserID, &item.TariffID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription by ID: %w", err)
	}
	return &item, nil
}

func (s sqlRepo) UpdateByID(ctx context.Context, item *Subscription) (*Subscription, error) {
	query := `UPDATE subscriptions SET user_id = $1, tariff_id = $2 WHERE id = $3 RETURNING id, user_id, tariff_id`
	row := s.db.QueryRow(ctx, query, item.UserID, item.TariffID, item.ID)
	var result Subscription
	err := row.Scan(&result.ID, &result.UserID, &result.TariffID)
	if err != nil {
		return nil, fmt.Errorf("failed to update subscription: %w", err)
	}
	return &result, nil
}

func (s sqlRepo) DeleteByID(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `DELETE FROM subscriptions WHERE id = $1`
	cmdTag, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete subscription: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return false, nil
	}
	return true, nil
}

func (s sqlRepo) List(ctx context.Context) ([]Subscription, error) {
	query := `SELECT id, user_id, tariff_id FROM subscriptions`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list subscriptions: %w", err)
	}
	defer rows.Close()

	var items []Subscription
	for rows.Next() {
		var item Subscription
		if err := rows.Scan(&item.ID, &item.UserID, &item.TariffID); err != nil {
			return nil, fmt.Errorf("failed to scan subscription: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

func (s sqlRepo) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM subscriptions WHERE id = $1)`
	var exists bool
	err := s.db.QueryRow(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if subscription exists: %w", err)
	}
	return exists, nil
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &sqlRepo{db}
}
