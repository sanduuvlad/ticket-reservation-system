package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/config"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/database"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/handler"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/repository"
	"github.com/sanduuvlad/ticket-reservation-system/services/event-service/internal/usecase"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool, err := database.NewPool(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// Repository
	venueRepository := repository.NewVenueRepository(pool)

	// Usecase
	venueUsecase := usecase.NewVenueUsecase(venueRepository)

	// Handler
	venueHandler := handler.NewVenueHandler(venueUsecase)

	// Router
	router := gin.Default()

	router.POST("/venues", venueHandler.CreateVenue)

	// Server
	address := fmt.Sprintf(":%d", cfg.HTTPServer.Port)

	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
