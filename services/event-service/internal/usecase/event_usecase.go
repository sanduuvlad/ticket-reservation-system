package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/apperror"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/domain"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/repository"
)

type EventUsecase interface {
	CreateEvent(ctx context.Context, req dto.CreateEventRequest) (dto.CreateEventResponse, error)
	GetAllEvents(ctx context.Context) ([]dto.EventResponse, error)
	GetByIDEvent(ctx context.Context, eventID int64) (dto.EventResponse, error)
}

type eventUsecase struct {
	repo      repository.EventRepository
	repoVenue repository.VenueRepository
}

func NewEventUsecase(repo repository.EventRepository, repoVenue repository.VenueRepository) *eventUsecase {
	return &eventUsecase{
		repo:      repo,
		repoVenue: repoVenue,
	}
}

func generateSeats(rowsCount, seatsPerRow int) []domain.Seat {
	seats := make([]domain.Seat, 0, rowsCount*seatsPerRow)

	for i := 0; i < rowsCount; i++ {
		row := string(rune('A' + i))

		for j := 1; j <= seatsPerRow; j++ {
			seat := domain.Seat{
				Row:    row,
				Number: int64(j),
			}

			seats = append(seats, seat)
		}
	}

	return seats
}

func validateEventTime(startTime, endTime time.Time) error {
	if !endTime.After(startTime) {
		return fmt.Errorf("validate event time: %w", apperror.ErrInvalidEventTime)
	}

	return nil
}

func (uc *eventUsecase) CreateEvent(ctx context.Context, req dto.CreateEventRequest) (dto.CreateEventResponse, error) {
	if err := validateEventTime(req.StartTime, req.EndTime); err != nil {
		return dto.CreateEventResponse{}, err
	}

	if req.Status != domain.EventStatusDraft && req.Status != domain.EventStatusCancelled && req.Status != domain.EventStatusPublished {
		return dto.CreateEventResponse{}, fmt.Errorf("validate event status: %w", apperror.ErrInvalidEventStatus)
	}

	if _, err := uc.repoVenue.GetByID(ctx, req.VenueID); err != nil {
		return dto.CreateEventResponse{}, fmt.Errorf("validate event venue: %w", err)
	}

	data := dto.CreateEventData{
		Name:        req.Name,
		Description: req.Description,
		VenueID:     req.VenueID,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Status:      req.Status,
		RowsCount:   req.RowsCount,
		SeatsPerRow: req.SeatsPerRow,
	}

	seats := generateSeats(req.RowsCount, req.SeatsPerRow)

	event, err := uc.repo.CreateEvent(ctx, data, seats)
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

func (uc *eventUsecase) GetAllEvents(ctx context.Context) ([]dto.EventResponse, error) {
	events, err := uc.repo.GetAllEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed get all events: %w", err)
	}

	eventsDTO := make([]dto.EventResponse, 0)

	for _, value := range events {
		eventDTO := dto.EventResponse{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			StartTime:   value.StartTime,
			EndTime:     value.EndTime,
			Status:      value.Status,
		}

		eventsDTO = append(eventsDTO, eventDTO)
	}

	return eventsDTO, nil
}

func (uc *eventUsecase) GetByIDEvent(ctx context.Context, eventID int64) (dto.EventResponse, error) {
	event, err := uc.repo.GetByIDEvent(ctx, eventID)
	if err != nil {
		return dto.EventResponse{}, err
	}

	eventDTO := dto.EventResponse{
		ID:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Status:      event.Status,
	}

	return eventDTO, nil
}
