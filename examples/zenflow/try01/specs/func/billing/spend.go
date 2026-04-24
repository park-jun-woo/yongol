package billing

import "fmt"

// @func spend
// @error 402
// @description Deducts Amount from Current and returns NewBalance.

type SpendRequest struct {
	Current int64
	Amount  int64
}

type SpendResponse struct {
	NewBalance int64
}

func Spend(req SpendRequest) (SpendResponse, error) {
	if req.Current < req.Amount {
		return SpendResponse{}, fmt.Errorf("insufficient credits")
	}
	return SpendResponse{NewBalance: req.Current - req.Amount}, nil
}
