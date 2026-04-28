package billing

import "errors"

// @func checkCredits
// @description Returns ErrInsufficientCredits when the org has no credits left.
// @error 402

type CheckCreditsRequest struct {
	Balance int64
}

type CheckCreditsResponse struct{}

// ErrInsufficientCredits is the typed error returned when the balance is zero or negative.
var ErrInsufficientCredits = errors.New("insufficient credits")

func CheckCredits(req CheckCreditsRequest) (CheckCreditsResponse, error) {
	if req.Balance <= 0 {
		return CheckCreditsResponse{}, ErrInsufficientCredits
	}
	return CheckCreditsResponse{}, nil
}
