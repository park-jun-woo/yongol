package workflow

import (
	"encoding/json"
	"github.com/jackc/pgx/v5/pgtype"
)

// @func batchProcessActions
// @description Validate and serialize action items for batch processing

type ActionItemInput struct {
	ActionType    string `json:"action_type"`
	Config        string `json:"config"`
	SequenceOrder int64  `json:"sequence_order"`
}

type BatchProcessActionsRequest struct {
	WorkflowID pgtype.UUID
	Items      string
}

type BatchProcessActionsResponse struct {
	ItemsJSON []byte
}

func BatchProcessActions(req BatchProcessActionsRequest) (BatchProcessActionsResponse, error) {
	// ff:what Validate and serialize action items for batch processing
	var items []ActionItemInput
	if err := json.Unmarshal([]byte(req.Items), &items); err != nil {
		return BatchProcessActionsResponse{}, err
	}
	out, err := json.Marshal(items)
	if err != nil {
		return BatchProcessActionsResponse{}, err
	}
	return BatchProcessActionsResponse{ItemsJSON: out}, nil
}
