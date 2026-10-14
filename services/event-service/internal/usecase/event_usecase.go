package usecase

import (
	"context"
	"fmt"

	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/repository"
)

type EventUsecase interface {
	CreateEvent(ctx context.Context, req dto.CreateEventRequest) (dto.CreateEventResponse, error)
}

type eventUsecase struct {
	repo repository.EventRepository
}

func NewEventUsecase(repo repository.EventRepository) *eventUsecase {
	return &eventUsecase{
		repo: repo,
	}
}

func (uc *eventUsecase) CreateEvent(ctx context.Context, req dto.CreateEventRequest) (dto.CreateEventResponse, error) {
	data := dto.CreateEventData{
		Name:        req.Name,
		Description: req.Description,
		VenueID:     req.VenueID,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Status:      req.Status,
	}

	event, err := uc.repo.CreateEvent(ctx, data)
	if err != nil {
		return dto.CreateEventResponse{}, fmt.Errorf("create event: %w", err)
	}

	eventDTO := dto.CreateEventResponse{
		ID:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Status:      event.Status,
	}

	return eventDTO, nil
}
