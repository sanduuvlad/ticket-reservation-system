package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/usecase"
)

type SeatHandler struct {
	usecase usecase.SeatUsecase
}

func NewSeatHandler(uc usecase.SeatUsecase) *SeatHandler {
	return &SeatHandler{
		usecase: uc,
	}
}

func (h *SeatHandler) GetByEventID(ctx *gin.Context) {
	eventIDStr := ctx.Param("id")

	eventID, err := strconv.ParseInt(eventIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	seatsResponse, err := h.usecase.GetByEventID(ctx.Request.Context(), eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	ctx.JSON(http.StatusOK, seatsResponse)
}
