package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/apperror"
)

func TestValidateEventTime(t *testing.T) {
	startTime := time.Date(2026, time.November, 13, 18, 0, 0, 0, time.UTC)
	endTime := time.Date(2026, time.November, 13, 20, 0, 0, 0, time.UTC)

	if err := validateEventTime(startTime, endTime); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	endTime = time.Date(2026, time.November, 13, 16, 0, 0, 0, time.UTC)

	if err := validateEventTime(startTime, endTime); !errors.Is(err, apperror.ErrInvalidEventTime) {
		t.Errorf("expected ErrInvalidEventTime, got: %v", err)
	}

	endTime = time.Date(2026, time.November, 13, 18, 0, 0, 0, time.UTC)

	if err := validateEventTime(startTime, endTime); !errors.Is(err, apperror.ErrInvalidEventTime) {
		t.Errorf("expected ErrInvalidEventTime, got: %v", err)
	}
}
