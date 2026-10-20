package usecase

import (
	"context"

	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/repository"
)

type SeatUsecase interface {
	GetByEventID(ctx context.Context, eventID int64) ([]dto.SeatResponse, error)
}

type seatUsecase struct {
	repo repository.SeatRepository
}

func NewSeatUsecase(repo repository.SeatRepository) *seatUsecase {
	return &seatUsecase{
		repo: repo,
	}
}

func (uc *seatUsecase) GetByEventID(ctx context.Context, eventID int64) ([]dto.SeatResponse, error) {
	seats, err := uc.repo.GetByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	seatsDTO := make([]dto.SeatResponse, 0)

	for _, seat := range seats {
		seatDTO := dto.SeatResponse{
			ID:     seat.ID,
			Row:    seat.Row,
			Number: seat.Number,
		}

		seatsDTO = append(seatsDTO, seatDTO)
	}

	return seatsDTO, nil
}
