package billing

// @func deductCredit
// @description Deducts credits from the organization (pure: returns the spent amount; persistence is handled by the SSaC layer).

type DeductCreditRequest struct {
	OrgID  int64
	Amount int64
}

type DeductCreditResponse struct {
	Spent int64
}

func DeductCredit(req DeductCreditRequest) (DeductCreditResponse, error) {
	return DeductCreditResponse{Spent: req.Amount}, nil
}
