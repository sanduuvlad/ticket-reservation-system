package apperror

import "errors"

var ErrEventNotFound = errors.New("event not found")

var ErrVenueNotFound = errors.New("venue not found")

var ErrInvalidEventTime = errors.New("end time must be after start time")

var ErrInvalidEventStatus = errors.New("invalid event status")
