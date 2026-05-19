package dashboard

// @func summarize
// @description Assemble org dashboard summary in-memory.

type SummarizeRequest struct {
	OrgName        string
	PlanType       string
	CreditsBalance int64
}

type SummarizeResponse struct {
	OrgName        string
	PlanType       string
	CreditsBalance int64
}

func Summarize(req SummarizeRequest) (SummarizeResponse, error) {
	return SummarizeResponse{
		OrgName:        req.OrgName,
		PlanType:       req.PlanType,
		CreditsBalance: req.CreditsBalance,
	}, nil
}
