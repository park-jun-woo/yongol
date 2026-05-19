package geocode

// @func verifyAddress
// @error 500
// @description Geocode an address via external geocoding API

type VerifyAddressRequest struct {
	Address string
}

type VerifyAddressResponse struct {
	Verified         bool
	FormattedAddress string
}

func VerifyAddress(req VerifyAddressRequest) (VerifyAddressResponse, error) {
	// Simulate geocoding API call
	return VerifyAddressResponse{
		Verified:         true,
		FormattedAddress: req.Address,
	}, nil
}
