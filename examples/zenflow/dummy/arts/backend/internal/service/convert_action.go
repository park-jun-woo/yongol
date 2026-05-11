//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertAction — db.Action row → *api.Action 변환
//ff:checked llm=yongol-gen hash=8e2b8ac9
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
)

func convertAction(row db.Action) (*api.Action, error) {
	return &api.Action{
		ActionType: ptrOf(row.ActionType),
		Config: ptrOf(row.Config),
		Id: ptrOf(row.ID),
		SequenceOrder: ptrOf(row.SequenceOrder),
		WorkflowId: ptrOf(row.WorkflowID),
	}, nil
}
