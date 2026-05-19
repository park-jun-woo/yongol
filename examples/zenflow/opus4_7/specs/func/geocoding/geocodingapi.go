package geocoding

// @func geocode
// @description Geocode an address and return coordinates. Stub for external API.

type GeocodeRequest struct {
	Address string
}

type GeocodeResponse struct {
	Latitude  string
	Longitude string
	Verified  bool
}

func Geocode(req GeocodeRequest) (GeocodeResponse, error) {
	return GeocodeResponse{
		Latitude:  "37.5665",
		Longitude: "126.9780",
		Verified:  true,
	}, nil
}
