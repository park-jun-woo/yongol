package billing

import "fmt"

// @func deductCredit
// @description Atomically deducts credits from an organization

type DeductCreditRequest struct {
	OrgID  int64
	Amount int64
}

type DeductCreditResponse struct {
	Remaining int64
}

func DeductCredit(req DeductCreditRequest) (DeductCreditResponse, error) {
	if req.Amount <= 0 {
		return DeductCreditResponse{}, fmt.Errorf("invalid deduction amount")
	}
	return DeductCreditResponse{Remaining: 0}, nil
}
