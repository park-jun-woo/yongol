package workflow

import "encoding/json"

// @func serializeActions
// @description Serialize action inputs to JSON for batch insert

type ActionInput struct {
	Type          string `json:"type"`
	Config        string `json:"config"`
	SequenceOrder int64  `json:"sequence_order"`
}

type SerializeActionsRequest struct {
	Items string
}

type SerializeActionsResponse struct {
	ItemsJSON []byte
}

func SerializeActions(req SerializeActionsRequest) (SerializeActionsResponse, error) {
	var items []ActionInput
	if err := json.Unmarshal([]byte(req.Items), &items); err != nil {
		return SerializeActionsResponse{}, err
	}
	out, err := json.Marshal(items)
	if err != nil {
		return SerializeActionsResponse{}, err
	}
	return SerializeActionsResponse{ItemsJSON: out}, nil
}
