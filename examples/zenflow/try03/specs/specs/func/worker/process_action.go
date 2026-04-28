package worker

import "fmt"

// @func processAction
// @description Simulates processing workflow actions via external API calls

type ProcessActionRequest struct {
	WorkflowID int64
}

type ProcessActionResponse struct {
	Processed int64
}

func ProcessAction(req ProcessActionRequest) (ProcessActionResponse, error) {
	if req.WorkflowID <= 0 {
		return ProcessActionResponse{}, fmt.Errorf("workflow ID required")
	}
	return ProcessActionResponse{Processed: 1}, nil
}
