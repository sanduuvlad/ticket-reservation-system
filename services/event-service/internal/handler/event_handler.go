package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/apperror"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/usecase"
)

type EventHandler struct {
	usecase usecase.EventUsecase
}

func NewEventHandler(uc usecase.EventUsecase) *EventHandler {
	return &EventHandler{
		usecase: uc,
	}
}

func (h *EventHandler) CreateEvent(ctx *gin.Context) {
	var request dto.CreateEventRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	eventResponse, err := h.usecase.CreateEvent(ctx.Request.Context(), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, eventResponse)
}

func (h *EventHandler) GetAllEvents(ctx *gin.Context) {
	eventsResponse, err := h.usecase.GetAllEvents(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	ctx.JSON(http.StatusOK, eventsResponse)
}

func (h *EventHandler) GetByIDEvent(ctx *gin.Context) {
	eventIDStr := ctx.Param("id")

	eventID, err := strconv.ParseInt(eventIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	eventResponse, err := h.usecase.GetByIDEvent(ctx.Request.Context(), eventID)
	if err != nil {
		if errors.Is(err, apperror.ErrEventNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "event not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	ctx.JSON(http.StatusOK, eventResponse)
}
