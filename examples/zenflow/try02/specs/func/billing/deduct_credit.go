package billing

import "github.com/jackc/pgx/v5/pgtype"

// @func deductCredit
// @error 500
// @description Atomically deducts credits from an organization's balance

type DeductCreditRequest struct {
	OrgID  pgtype.UUID
	Amount int64
}

type DeductCreditResponse struct {
	RemainingBalance int64
}

func DeductCredit(req DeductCreditRequest) (DeductCreditResponse, error) {
	// In generated code, this would atomically update the DB
	return DeductCreditResponse{RemainingBalance: 0}, nil
}
