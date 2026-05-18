package schedule

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// @func scheduleKey
// @error 500
// @description Converts a workflow UUID to a string session key for schedule storage.

type ScheduleKeyRequest struct {
	WorkflowID openapi_types.UUID
}

type ScheduleKeyResponse struct {
	Key string
}

func ScheduleKey(req ScheduleKeyRequest) (ScheduleKeyResponse, error) {
	return ScheduleKeyResponse{Key: "schedule:" + req.WorkflowID.String()}, nil
}
