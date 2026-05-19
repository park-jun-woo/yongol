package geocoding

// @func geocode
// @description Simulate geocoding API call (purity-safe mock)

type GeocodeRequest struct {
	Address string
}

type GeocodeResponse struct {
	Latitude  float64
	Longitude float64
}

func Geocode(req GeocodeRequest) (GeocodeResponse, error) {
	// Simulated geocoding - purity rules forbid real HTTP
	return GeocodeResponse{Latitude: 37.5665, Longitude: 126.9780}, nil
}
