package worker

// @func processAction
// @description Simulates processing an external action

type ProcessActionRequest struct {
	ActionType string
	Config     string
}

type ProcessActionResponse struct {
	Success bool
}

func ProcessAction(req ProcessActionRequest) (ProcessActionResponse, error) {
	return ProcessActionResponse{Success: true}, nil
}
