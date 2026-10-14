package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
