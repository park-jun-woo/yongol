package schedule

import (
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// @func scheduleKey
// @description Generate session key for workflow schedule

type ScheduleKeyRequest struct {
	WorkflowID openapi_types.UUID
}

type ScheduleKeyResponse struct {
	Key string
}

func ScheduleKey(req ScheduleKeyRequest) (ScheduleKeyResponse, error) {
	key := fmt.Sprintf("schedule:%s", req.WorkflowID.String())
	return ScheduleKeyResponse{Key: key}, nil
}
