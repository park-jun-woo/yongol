package schedule

import openapi_types "github.com/oapi-codegen/runtime/types"

// @func buildKey
// @description Build a session key from workflow ID

type BuildKeyRequest struct {
	WorkflowID openapi_types.UUID
}

type BuildKeyResponse struct {
	Key string
}

func BuildKey(req BuildKeyRequest) (BuildKeyResponse, error) {
	return BuildKeyResponse{Key: "schedule:" + req.WorkflowID.String()}, nil
}
