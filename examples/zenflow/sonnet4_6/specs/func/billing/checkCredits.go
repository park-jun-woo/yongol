package billing

// @func checkCredits
// @error 402
// @description Returns true if the organization balance is zero or below (triggers 402)

type CheckCreditsRequest struct {
	Balance int64
}

// Returns true = insufficient credits (will trigger 402 error)
func CheckCredits(req CheckCreditsRequest) bool {
	return req.Balance <= 0
}
