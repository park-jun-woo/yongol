//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertAction — db.Action row → *api.Action 변환
//ff:checked llm=yongol-gen hash=c6bbd808
package service

import (
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
)

func convertAction(row db.Action) (*api.Action, error) {
	return &api.Action{
		ActionType: ptrOf(row.ActionType),
		Id: ptrOf(row.ID),
		PayloadTemplate: ptrOf(row.PayloadTemplate),
		SequenceOrder: ptrOf(row.SequenceOrder),
		WorkflowId: ptrOf(row.WorkflowID),
	}, nil
}
