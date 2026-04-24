//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertExecutionLog — db.ExecutionLog row → *api.ExecutionLog 변환
//ff:checked llm=yongol-gen hash=48f13eee
package service

import (
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
)

func convertExecutionLog(row db.ExecutionLog) (*api.ExecutionLog, error) {
	return &api.ExecutionLog{
		CreditsSpent: ptrOf(row.CreditsSpent),
		ExecutedAt: ptrOf(row.ExecutedAt.Time),
		Id: ptrOf(row.ID),
		OrgId: ptrOf(row.OrgID),
		Status: ptrOf(row.Status),
		WorkflowId: ptrOf(row.WorkflowID),
	}, nil
}
