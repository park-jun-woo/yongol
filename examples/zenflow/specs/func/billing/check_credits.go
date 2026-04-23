package billing

import "fmt"

// @func checkCredits
// @error 402
// @description Fail with 402 when current balance cannot afford the amount.

type CheckCreditsRequest struct {
	Current int64
	Amount  int64
}

type CheckCreditsResponse struct {
	OK bool
}

func CheckCredits(req CheckCreditsRequest) (CheckCreditsResponse, error) {
	if req.Current < req.Amount {
		return CheckCreditsResponse{}, fmt.Errorf("insufficient credits")
	}
	return CheckCreditsResponse{OK: true}, nil
}
