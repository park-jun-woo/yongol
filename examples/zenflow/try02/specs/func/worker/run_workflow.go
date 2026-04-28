package worker

// @func runWorkflow
// @description Executes the workflow's actions (simulated; pure — no IO).

type RunWorkflowRequest struct {
	WorkflowID int64
}

type RunWorkflowResponse struct {
	Status string
}

func RunWorkflow(req RunWorkflowRequest) (RunWorkflowResponse, error) {
	_ = req.WorkflowID
	return RunWorkflowResponse{Status: "success"}, nil
}
