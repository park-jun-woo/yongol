package geocoding

import (
	"context"
	"os"
)

// @func geocode
// @description Call external geocoding API to verify an address

type GeocodeRequest struct {
	Address string
}

// Response type is already defined in geocodingapi.go as GeocodeResponse

func Geocode(req GeocodeRequest) (GeocodeResponse, error) {
	baseURL := os.Getenv("GEOCODING_API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:9999"
	}
	client := NewGeocodingAPIModel(baseURL)
	resp, err := client.Geocode(context.Background(), req.Address)
	if err != nil {
		return GeocodeResponse{}, err
	}
	return *resp, nil
}
