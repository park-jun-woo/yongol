package schedule

import openapi_types "github.com/oapi-codegen/runtime/types"

// @func buildScheduleKey
// @description Build a session key from workflow ID

type BuildScheduleKeyRequest struct {
	WorkflowID openapi_types.UUID
}

type BuildScheduleKeyResponse struct {
	Key string
}

func BuildScheduleKey(req BuildScheduleKeyRequest) (BuildScheduleKeyResponse, error) {
	return BuildScheduleKeyResponse{Key: "schedule:" + req.WorkflowID.String()}, nil
}
