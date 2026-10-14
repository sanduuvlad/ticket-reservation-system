package usecase

import (
	"context"
	"fmt"

	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/repository"
)

type VenueUsecase interface {
	CreateVenue(ctx context.Context, req dto.CreateVenueRequest) (dto.CreateVenueResponse, error)
}

type venueUsecase struct {
	repo repository.VenueRepository
}

func NewVenueUsecase(repo repository.VenueRepository) *venueUsecase {
	return &venueUsecase{
		repo: repo,
	}
}

func (uc *venueUsecase) CreateVenue(ctx context.Context, req dto.CreateVenueRequest) (dto.CreateVenueResponse, error) {
	data := dto.CreateVenueData{
		Name:    req.Name,
		Country: req.Country,
		City:    req.City,
		Address: req.Address,
	}

	venue, err := uc.repo.CreateVenue(ctx, data)
	if err != nil {
		return dto.CreateVenueResponse{}, fmt.Errorf("create venue: %w", err)
	}

	venueDTO := dto.CreateVenueResponse{
		ID:      venue.ID,
		Name:    venue.Name,
		Country: venue.Country,
		City:    venue.City,
		Address: venue.Address,
	}

	return venueDTO, nil
}
