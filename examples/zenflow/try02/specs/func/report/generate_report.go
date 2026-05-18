package report

import (
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// @func generateReport
// @error 500
// @description Formats workflow execution result as a text report (purity-safe) and returns a storage key

type GenerateReportRequest struct {
	WorkflowID  openapi_types.UUID
	ActionCount int64
	Status      string
}

type GenerateReportResponse struct {
	Key     string
	Content string
}

func GenerateReport(req GenerateReportRequest) (GenerateReportResponse, error) {
	key := fmt.Sprintf("reports/%s.txt", req.WorkflowID.String())
	content := fmt.Sprintf(
		"Execution Report\nWorkflow: %s\nActions: %d\nStatus: %s\n",
		req.WorkflowID.String(), req.ActionCount, req.Status,
	)
	return GenerateReportResponse{Key: key, Content: content}, nil
}
