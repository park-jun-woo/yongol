package dashboard

// @func buildExecutionDetail
// @description Compose ExecutionDetail from log + workflow + org

type BuildExecutionDetailRequest struct {
	LogID         string
	WorkflowID    string
	WorkflowTitle string
	OrgID         string
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
		ID:            req.LogID,
		WorkflowID:    req.WorkflowID,
		WorkflowTitle: req.WorkflowTitle,
		OrgID:         req.OrgID,
		OrgName:       req.OrgName,
		Status:        req.Status,
		CreditsSpent:  req.CreditsSpent,
		ExecutedAt:    req.ExecutedAt,
	}, nil
}
