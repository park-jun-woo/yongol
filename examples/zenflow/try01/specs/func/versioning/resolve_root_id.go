package versioning

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// @func resolveRootID
// @description If root_workflow_id is zero UUID, return the workflow's own ID; else return existing root

type ResolveRootIDRequest struct {
	WorkflowID       pgtype.UUID
	RootWorkflowID   pgtype.UUID
}

type ResolveRootIDResponse struct {
	RootID pgtype.UUID
}

func ResolveRootID(req ResolveRootIDRequest) (ResolveRootIDResponse, error) {
	zeroUUID := pgtype.UUID{}
	if req.RootWorkflowID == zeroUUID || !req.RootWorkflowID.Valid {
		return ResolveRootIDResponse{RootID: req.WorkflowID}, nil
	}
	return ResolveRootIDResponse{RootID: req.RootWorkflowID}, nil
}
