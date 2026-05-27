package version

// @func resolveRootID
// @description If root_workflow_id is zero, return own ID; else return existing root

type ResolveRootIDRequest struct {
	WorkflowID     int64
	RootWorkflowID int64
}

type ResolveRootIDResponse struct {
	RootID int64
}

func ResolveRootID(req ResolveRootIDRequest) (ResolveRootIDResponse, error) {
	if req.RootWorkflowID == 0 {
		return ResolveRootIDResponse{RootID: req.WorkflowID}, nil
	}
	return ResolveRootIDResponse{RootID: req.RootWorkflowID}, nil
}
