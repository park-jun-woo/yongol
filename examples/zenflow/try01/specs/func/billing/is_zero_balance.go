package billing

// @func isZeroBalance
// @description Check if credits balance is zero or below

type IsZeroBalanceRequest struct {
	Balance int64
}

func IsZeroBalance(req IsZeroBalanceRequest) bool {
	return req.Balance <= 0
}
