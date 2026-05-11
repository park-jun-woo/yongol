//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertWorkflow — db.Workflow row → *api.Workflow 변환
//ff:checked llm=yongol-gen hash=28b71853
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/ssac/pkg/pgtypex"
)

func convertWorkflow(row db.Workflow) (*api.Workflow, error) {
	return &api.Workflow{
		CreatedAt: ptrOf(pgtypex.FromPgTimestamptz(row.CreatedAt)),
		Id: ptrOf(row.ID),
		OrgId: ptrOf(row.OrgID),
		RootWorkflowId: ptrOf(row.RootWorkflowID),
		Status: ptrOf(row.Status),
		Title: ptrOf(row.Title),
		TriggerEvent: ptrOf(row.TriggerEvent),
		Version: ptrOf(row.Version),
	}, nil
}
