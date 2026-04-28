package billing

import "fmt"

// @func checkCredits
// @error 402
// @description Returns error when the organization has zero or negative credits balance

type CheckCreditsRequest struct {
	Balance int64
}

type CheckCreditsResponse struct {
	Sufficient bool
}

func CheckCredits(req CheckCreditsRequest) (CheckCreditsResponse, error) {
	if req.Balance <= 0 {
		return CheckCreditsResponse{Sufficient: false}, fmt.Errorf("insufficient credits")
	}
	return CheckCreditsResponse{Sufficient: true}, nil
}
