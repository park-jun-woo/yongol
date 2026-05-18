package workflow

import (
	"encoding/json"

	"github.com/park-jun-woo/zenflow/internal/api"
)

// @func serializeActions
// @error 500
// @description Serializes a slice of ActionInput to JSON bytes for batch INSERT via jsonb_array_elements.

type SerializeActionsRequest struct {
	Actions []api.ActionInput
}

type SerializeActionsResponse struct {
	ItemsJSON []byte
	Count     int64
}

func SerializeActions(req SerializeActionsRequest) (SerializeActionsResponse, error) {
	b, err := json.Marshal(req.Actions)
	if err != nil {
		return SerializeActionsResponse{}, err
	}
	return SerializeActionsResponse{ItemsJSON: b, Count: int64(len(req.Actions))}, nil
}
