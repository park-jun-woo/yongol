package billing

// @func isInsufficient
// @error 402
// @description Returns true if credits balance is insufficient (zero or below)

type IsInsufficientRequest struct {
	Balance int64
}

func IsInsufficient(req IsInsufficientRequest) bool {
	return req.Balance <= 0
}
