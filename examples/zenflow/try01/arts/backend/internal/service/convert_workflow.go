//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertWorkflow — db.Workflow row → *api.Workflow 변환
//ff:checked llm=yongol-gen hash=63a9fa79
package service

import (
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
)

func convertWorkflow(row db.Workflow) (*api.Workflow, error) {
	return &api.Workflow{
		CreatedAt: ptrOf(row.CreatedAt),
		Id: ptrOf(row.ID),
		OrgId: ptrOf(row.OrgID),
		Status: ptrOf(row.Status),
		Title: ptrOf(row.Title),
		TriggerEvent: ptrOf(row.TriggerEvent),
	}, nil
}
