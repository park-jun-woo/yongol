package billing

import "fmt"

// @func checkCredits
// @error 402
// @description Fails when current balance is not positive.

type CheckCreditsRequest struct {
	Current int64
}

type CheckCreditsResponse struct {
	OK bool
}

func CheckCredits(req CheckCreditsRequest) (CheckCreditsResponse, error) {
	if req.Current <= 0 {
		return CheckCreditsResponse{}, fmt.Errorf("insufficient credits")
	}
	return CheckCreditsResponse{OK: true}, nil
}
