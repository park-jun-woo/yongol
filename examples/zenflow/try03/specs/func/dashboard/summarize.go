package dashboard

import "github.com/jackc/pgx/v5/pgtype"

// @func summarize
// @description Assemble org dashboard summary in-memory.

type SummarizeRequest struct {
	OrgID          pgtype.UUID
	OrgName        string
	PlanType       string
	CreditsBalance int64
}

type SummarizeResponse struct {
	OrgName           string
	PlanType          string
	CreditsBalance    int64
	ActiveWorkflows   int64
	PausedWorkflows   int64
	TotalExecutions   int64
	TotalCreditsSpent int64
}

func Summarize(req SummarizeRequest) (SummarizeResponse, error) {
	return SummarizeResponse{
		OrgName:           req.OrgName,
		PlanType:          req.PlanType,
		CreditsBalance:    req.CreditsBalance,
		ActiveWorkflows:   3,
		PausedWorkflows:   1,
		TotalExecutions:   10,
		TotalCreditsSpent: 5,
	}, nil
}
