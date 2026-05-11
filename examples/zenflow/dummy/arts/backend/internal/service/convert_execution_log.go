//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertExecutionLog — db.ExecutionLog row → *api.ExecutionLog 변환
//ff:checked llm=yongol-gen hash=e28750ed
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/ssac/pkg/pgtypex"
)

func convertExecutionLog(row db.ExecutionLog) (*api.ExecutionLog, error) {
	return &api.ExecutionLog{
		CreditsSpent: ptrOf(row.CreditsSpent),
		ExecutedAt: ptrOf(pgtypex.FromPgTimestamptz(row.ExecutedAt)),
		Id: ptrOf(row.ID),
		OrgId: ptrOf(row.OrgID),
		Status: ptrOf(row.Status),
		WorkflowId: ptrOf(row.WorkflowID),
	}, nil
}
