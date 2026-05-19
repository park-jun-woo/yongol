package report

import (
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// @func generateReport
// @description Format execution result as text report

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
	key := fmt.Sprintf("reports/%s/report.txt", req.WorkflowID.String())
	content := fmt.Sprintf("Workflow: %s\nActions: %d\nStatus: %s\n", req.WorkflowID.String(), req.ActionCount, req.Status)
	return GenerateReportResponse{Key: key, Content: content}, nil
}
