package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/domain"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
)

type VenueRepository interface {
	CreateVenue(ctx context.Context, req dto.CreateVenueData) (domain.Venue, error)
}

type PostgresVenueRepository struct {
	pool *pgxpool.Pool
}

func NewVenueRepository(pool *pgxpool.Pool) *PostgresVenueRepository {
	return &PostgresVenueRepository{
		pool: pool,
	}
}

func (r *PostgresVenueRepository) CreateVenue(ctx context.Context, req dto.CreateVenueData) (domain.Venue, error) {
	const query = `
		INSERT INTO venues (name, country, city, address)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, country, city, address
	`

	var venue domain.Venue

	err := r.pool.QueryRow(
		ctx,
		query,
		req.Name,
		req.Country,
		req.City,
		req.Address,
	).Scan(
		&venue.ID,
		&venue.Name,
		&venue.Country,
		&venue.City,
		&venue.Address,
	)

	if err != nil {
		return domain.Venue{}, fmt.Errorf("create venue: %w", err)
	}

	return venue, nil
}
