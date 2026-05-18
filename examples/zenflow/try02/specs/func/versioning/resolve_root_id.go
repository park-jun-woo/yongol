package versioning

// @func resolveRootID
// @error 500
// @description Returns the root workflow ID. If root_workflow_id is the zero UUID, returns the workflow's own ID; otherwise returns the existing root.

import "github.com/jackc/pgx/v5/pgtype"

type ResolveRootIDRequest struct {
	WorkflowID     pgtype.UUID
	RootWorkflowID pgtype.UUID
}

type ResolveRootIDResponse struct {
	RootID pgtype.UUID
}

var zeroBytes [16]byte

func ResolveRootID(req ResolveRootIDRequest) (ResolveRootIDResponse, error) {
	if req.RootWorkflowID.Bytes == zeroBytes {
		return ResolveRootIDResponse{RootID: req.WorkflowID}, nil
	}
	return ResolveRootIDResponse{RootID: req.RootWorkflowID}, nil
}
