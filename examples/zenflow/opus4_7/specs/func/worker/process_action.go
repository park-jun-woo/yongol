package worker

// @func processAction
// @description Simulates processing a workflow action (external API call stub)

type ProcessActionRequest struct {
	ActionType string
	Config     string
}

type ProcessActionResponse struct {
	Success bool
}

func ProcessAction(req ProcessActionRequest) (ProcessActionResponse, error) {
	return ProcessActionResponse{
		Success: true,
	}, nil
}
