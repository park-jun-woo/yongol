//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertWorkflow — db.Workflow row → *api.Workflow 변환
//ff:checked llm=yongol-gen hash=e168a818
package service

import (
	"github.com/example/zenflow_try01/internal/api"
	"github.com/example/zenflow_try01/internal/db"
)

func convertWorkflow(row db.Workflow) (*api.Workflow, error) {
	return &api.Workflow{
		CreatedAt: row.CreatedAt.Time,
		Id: row.ID,
		OwnerId: row.OwnerID,
		Status: api.WorkflowStatus(row.Status),
		Title: row.Title,
		TriggerEvent: row.TriggerEvent,
	}, nil
}
