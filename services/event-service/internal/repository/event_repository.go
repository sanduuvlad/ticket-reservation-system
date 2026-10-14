package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/domain"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, req dto.CreateEventData) (domain.Event, error)
}

type PostgresEventRepository struct {
	pool *pgxpool.Pool
}

func NewEventRepository(pool *pgxpool.Pool) *PostgresEventRepository {
	return &PostgresEventRepository{
		pool: pool,
	}
}

func (r *PostgresEventRepository) CreateEvent(ctx context.Context, req dto.CreateEventData) (domain.Event, error) {
	const query = `
		INSERT INTO events (name, description, venue_id,
			start_time, end_time, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, description, venue_id,
			start_time, end_time, status
	`

	var event domain.Event

	err := r.pool.QueryRow(
		ctx,
		query,
		req.Name,
		req.Description,
		req.VenueID,
		req.StartTime,
		req.EndTime,
		req.Status,
	).Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.VenueID,
		&event.StartTime,
		&event.EndTime,
		&event.Status,
	)

	if err != nil {
		return domain.Event{}, fmt.Errorf("create event: %w", err)
	}

	return event, nil
}
