package billing

import "errors"

// @func checkCredits
// @error 402
// @description Returns OK when balance > 0, otherwise an error so @call surfaces a 402.

type CheckCreditsRequest struct {
	Balance int64
}

type CheckCreditsResponse struct {
}

func CheckCredits(req CheckCreditsRequest) (CheckCreditsResponse, error) {
	if req.Balance <= 0 {
		return CheckCreditsResponse{}, errors.New("insufficient credits")
	}
	return CheckCreditsResponse{}, nil
}
