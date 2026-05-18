package dashboard

import "github.com/jackc/pgx/v5/pgtype"

// @func buildExecutionDetail
// @description Compose ExecutionDetail from log + workflow + org.

type BuildExecutionDetailRequest struct {
	LogID         pgtype.UUID
	WorkflowID    pgtype.UUID
	WorkflowTitle string
	OrgID         pgtype.UUID
	OrgName       string
	Status        string
	CreditsSpent  int64
	ExecutedAt    string
}

type BuildExecutionDetailResponse struct {
	ID            string
	WorkflowID    string
	WorkflowTitle string
	OrgID         string
	OrgName       string
	Status        string
	CreditsSpent  int64
	ExecutedAt    string
}

func BuildExecutionDetail(req BuildExecutionDetailRequest) (BuildExecutionDetailResponse, error) {
	return BuildExecutionDetailResponse{
		ID:            req.LogID.String(),
		WorkflowID:    req.WorkflowID.String(),
		WorkflowTitle: req.WorkflowTitle,
		OrgID:         req.OrgID.String(),
		OrgName:       req.OrgName,
		Status:        req.Status,
		CreditsSpent:  req.CreditsSpent,
		ExecutedAt:    req.ExecutedAt,
	}, nil
}
