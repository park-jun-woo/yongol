package dashboard

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

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
	LogID         string
	WorkflowID    string
	WorkflowTitle string
	OrgID         string
	OrgName       string
	Status        string
	CreditsSpent  int64
	ExecutedAt    string
}

func uuidStr(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	b := u.Bytes
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func BuildExecutionDetail(req BuildExecutionDetailRequest) (BuildExecutionDetailResponse, error) {
	return BuildExecutionDetailResponse{
		LogID:         uuidStr(req.LogID),
		WorkflowID:    uuidStr(req.WorkflowID),
		WorkflowTitle: req.WorkflowTitle,
		OrgID:         uuidStr(req.OrgID),
		OrgName:       req.OrgName,
		Status:        req.Status,
		CreditsSpent:  req.CreditsSpent,
		ExecutedAt:    req.ExecutedAt,
	}, nil
}
