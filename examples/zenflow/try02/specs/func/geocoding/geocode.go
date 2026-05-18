package geocoding

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
)

// @func geocode
// @error 500
// @description Geocodes an address using the external geocoding API client.

type GeocodeRequest struct {
	Address string
}

type GeocodeResponse struct {
	Latitude  pgtype.Float8
	Longitude pgtype.Float8
}

func Geocode(req GeocodeRequest) (GeocodeResponse, error) {
	baseURL := os.Getenv("GEOCODING_API_URL")
	if baseURL == "" {
		baseURL = "https://api.geocoding.example.com"
	}
	model := NewGeocodingModel(baseURL)
	result, err := model.Geocode(context.Background(), req.Address)
	if err != nil {
		return GeocodeResponse{}, fmt.Errorf("geocode: %w", err)
	}
	return GeocodeResponse{
		Latitude:  pgtype.Float8{Float64: float64(result.Latitude), Valid: true},
		Longitude: pgtype.Float8{Float64: float64(result.Longitude), Valid: true},
	}, nil
}
