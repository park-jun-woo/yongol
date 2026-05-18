package worker

import "github.com/jackc/pgx/v5/pgtype"

// @func processActions
// @description Simulates processing all actions in a workflow

type ProcessActionsRequest struct {
	WorkflowID pgtype.UUID
}

type ProcessActionsResponse struct {
	Status string
}

func ProcessActions(req ProcessActionsRequest) (ProcessActionsResponse, error) {
	return ProcessActionsResponse{Status: "completed"}, nil
}
