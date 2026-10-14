package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/dto"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/usecase"
)

type VenueHandler struct {
	usecase usecase.VenueUsecase
}

func NewVenueHandler(uc usecase.VenueUsecase) *VenueHandler {
	return &VenueHandler{
		usecase: uc,
	}
}

func (h *VenueHandler) CreateVenue(ctx *gin.Context) {
	var request dto.CreateVenueRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	venueResponse, err := h.usecase.CreateVenue(ctx.Request.Context(), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, venueResponse)
}
