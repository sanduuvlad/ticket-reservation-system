package dto

type CreateVenueRequest struct {
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
	City    string `json:"city" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type CreateVenueResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
}

type CreateVenueData struct {
	Name    string
	Country string
	City    string
	Address string
}
