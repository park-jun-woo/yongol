package workflow

import (
	"encoding/json"

	"github.com/park-jun-woo/zenflow/internal/api"
)

// @func serializeActions
// @description Serialize action items array to JSON for batch insert

type SerializeActionsRequest struct {
	Actions []api.ActionInput
}

type SerializeActionsResponse struct {
	ItemsJSON []byte
}

func SerializeActions(req SerializeActionsRequest) (SerializeActionsResponse, error) {
	type item struct {
		Type          string `json:"type"`
		Config        string `json:"config"`
		SequenceOrder int64  `json:"sequence_order"`
	}
	items := make([]item, len(req.Actions))
	for i, a := range req.Actions {
		items[i] = item{
			Type:          a.ActionType,
			Config:        a.Config,
			SequenceOrder: a.SequenceOrder,
		}
	}
	data, err := json.Marshal(items)
	if err != nil {
		return SerializeActionsResponse{}, err
	}
	return SerializeActionsResponse{ItemsJSON: data}, nil
}
