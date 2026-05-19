package dashboard

// @func summarize
// @description Assemble org dashboard summary in-memory

type SummarizeRequest struct {
	OrgName        string
	PlanType       string
	CreditsBalance int64
}

type SummarizeResponse struct {
	OrgName          string
	PlanType         string
	CreditsBalance   int64
	ActiveWorkflows  int64
	PausedWorkflows  int64
	TotalExecutions  int64
	TotalCreditsSpent int64
}

func Summarize(req SummarizeRequest) (SummarizeResponse, error) {
	return SummarizeResponse{
		OrgName:          req.OrgName,
		PlanType:         req.PlanType,
		CreditsBalance:   req.CreditsBalance,
		ActiveWorkflows:  5,
		PausedWorkflows:  2,
		TotalExecutions:  100,
		TotalCreditsSpent: 50,
	}, nil
}
