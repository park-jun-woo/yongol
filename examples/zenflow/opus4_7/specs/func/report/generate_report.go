package report

import "fmt"

// @func generateReport
// @description Format execution result as text report

type GenerateReportRequest struct {
	WorkflowTitle string
	ActionCount   int64
	Status        string
}

type GenerateReportResponse struct {
	ReportKey  string
	ReportBody string
}

func GenerateReport(req GenerateReportRequest) (GenerateReportResponse, error) {
	key := fmt.Sprintf("reports/%s/%s", req.WorkflowTitle, req.Status)
	body := fmt.Sprintf("Workflow: %s\nActions: %d\nStatus: %s", req.WorkflowTitle, req.ActionCount, req.Status)
	return GenerateReportResponse{ReportKey: key, ReportBody: body}, nil
}
