package worker

// @func processActions
// @description Simulates processing all workflow actions in sequence

type ProcessActionsRequest struct {
	WorkflowID int64
	ActionCount int64
}

type ProcessActionsResponse struct {
	Status string
}

func ProcessActions(req ProcessActionsRequest) (ProcessActionsResponse, error) {
	return ProcessActionsResponse{Status: "success"}, nil
}
