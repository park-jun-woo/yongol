package report

import (
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// @func generateReport
// @description Format the execution result as text for file storage

type GenerateReportRequest struct {
	WorkflowID  openapi_types.UUID
	ActionCount int64
	Status      string
}

type GenerateReportResponse struct {
	Key  string
	Body string
}

func GenerateReport(req GenerateReportRequest) (GenerateReportResponse, error) {
	wfID := req.WorkflowID.String()
	key := fmt.Sprintf("reports/%s/%s.txt", wfID, req.Status)
	body := fmt.Sprintf("Workflow: %s\nActions: %d\nStatus: %s\n", wfID, req.ActionCount, req.Status)
	return GenerateReportResponse{Key: key, Body: body}, nil
}
