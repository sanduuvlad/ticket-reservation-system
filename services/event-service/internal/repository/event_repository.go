package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/apperror"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/domain"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, req dto.CreateEventData, seats []domain.Seat) (domain.Event, error)
	GetAllEvents(ctx context.Context) ([]domain.Event, error)
	GetByIDEvent(ctx context.Context, eventID int64) (domain.Event, error)
}

type PostgresEventRepository struct {
	pool *pgxpool.Pool
}

func NewEventRepository(pool *pgxpool.Pool) *PostgresEventRepository {
	return &PostgresEventRepository{
		pool: pool,
	}
}

func (r *PostgresEventRepository) CreateEvent(
	ctx context.Context,
	req dto.CreateEventData,
	seats []domain.Seat,
) (domain.Event, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Event{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	const query = `
		INSERT INTO events (name, description, venue_id,
			start_time, end_time, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, description, venue_id,
			start_time, end_time, status
	`

	var event domain.Event

	err = tx.QueryRow(
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

	const insertSeatQuery = `
		INSERT INTO seats (event_id, seat_row, seat_number)
		VALUES ($1, $2, $3)
	`
	for _, seat := range seats {

		_, err := tx.Exec(
			ctx,
			insertSeatQuery,
			event.ID,
			seat.Row,
			seat.Number,
		)
		if err != nil {
			return domain.Event{}, fmt.Errorf("create seat: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return domain.Event{}, fmt.Errorf("commit transaction: %w", err)
	}
	return event, nil
}

func (r *PostgresEventRepository) GetAllEvents(ctx context.Context) ([]domain.Event, error) {
	const query = `
		SELECT id, name, description, venue_id,
			start_time, end_time, status
		FROM events
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed get all events: %w", err)
	}
	defer rows.Close()

	var events []domain.Event

	for rows.Next() {
		var event domain.Event

		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.VenueID,
			&event.StartTime,
			&event.EndTime,
			&event.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to extract: %w", err)
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}

	return events, nil
}

func (r *PostgresEventRepository) GetByIDEvent(ctx context.Context, eventID int64) (domain.Event, error) {
	const query = `
		SELECT id, name, description, venue_id,
			start_time, end_time, status
		FROM events
		WHERE id = $1
	`

	var event domain.Event

	err := r.pool.QueryRow(
		ctx,
		query,
		eventID,
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
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Event{}, apperror.ErrEventNotFound
		}

		return domain.Event{}, fmt.Errorf("get event by id: %w", err)
	}

	return event, nil
}
