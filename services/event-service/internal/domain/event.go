package domain

import "time"

type EventStatus string

const (
	EventStatusDraft     EventStatus = "DRAFT"
	EventStatusPublished EventStatus = "PUBLISHED"
	EventStatusCancelled EventStatus = "CANCELLED"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	VenueID     int64
	StartTime   time.Time
	EndTime     time.Time
	Status      EventStatus
}
