package billing

import "fmt"

// @func checkCredits
// @error 402
// @description Returns error if credits balance is zero or negative

type CheckCreditsRequest struct {
	Balance int64
}

type CheckCreditsResponse struct{}

func CheckCredits(req CheckCreditsRequest) (CheckCreditsResponse, error) {
	if req.Balance <= 0 {
		return CheckCreditsResponse{}, fmt.Errorf("insufficient credits")
	}
	return CheckCreditsResponse{}, nil
}
