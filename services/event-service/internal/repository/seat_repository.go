package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/domain"
)

type SeatRepository interface {
	GetByEventID(ctx context.Context, eventID int64) ([]domain.Seat, error)
}

type PostgresSeatRepository struct {
	pool *pgxpool.Pool
}

func NewSeatRepository(pool *pgxpool.Pool) *PostgresSeatRepository {
	return &PostgresSeatRepository{
		pool: pool,
	}
}

func (r *PostgresSeatRepository) GetByEventID(ctx context.Context, eventID int64) ([]domain.Seat, error) {
	const query = `
		SELECT id, event_id, seat_row, seat_number
		FROM seats
		WHERE event_id = $1
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		eventID,
	)
	if err != nil {
		return nil, fmt.Errorf("get seats: %w", err)
	}
	defer rows.Close()

	var seats []domain.Seat

	for rows.Next() {
		var seat domain.Seat

		err := rows.Scan(
			&seat.ID,
			&seat.EventID,
			&seat.Row,
			&seat.Number,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan seat: %w", err)
		}

		seats = append(seats, seat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get seats: %w", err)
	}

	return seats, nil
}
