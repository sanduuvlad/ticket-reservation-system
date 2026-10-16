package dto

import (
	"time"

	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/domain"
)

type CreateEventRequest struct {
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description"`
	VenueID     int64              `json:"venue_id" binding:"required"`
	StartTime   time.Time          `json:"start_time" binding:"required"`
	EndTime     time.Time          `json:"end_time" binding:"required"`
	Status      domain.EventStatus `json:"status" binding:"required"`
}

type CreateEventResponse struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	StartTime   time.Time          `json:"start_time"`
	EndTime     time.Time          `json:"end_time"`
	Status      domain.EventStatus `json:"status"`
}

type CreateEventData struct {
	Name        string
	Description string
	VenueID     int64
	StartTime   time.Time
	EndTime     time.Time
	Status      domain.EventStatus
}

type EventResponse struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	StartTime   time.Time          `json:"start_time"`
	EndTime     time.Time          `json:"end_time"`
	Status      domain.EventStatus `json:"status"`
}
