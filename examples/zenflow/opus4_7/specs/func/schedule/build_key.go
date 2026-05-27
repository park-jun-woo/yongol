package schedule

import "fmt"

// @func buildKey
// @description Build a session key from workflow ID

type BuildKeyRequest struct {
	WorkflowID int64
}

type BuildKeyResponse struct {
	Key string
}

func BuildKey(req BuildKeyRequest) (BuildKeyResponse, error) {
	return BuildKeyResponse{Key: fmt.Sprintf("schedule:%d", req.WorkflowID)}, nil
}
