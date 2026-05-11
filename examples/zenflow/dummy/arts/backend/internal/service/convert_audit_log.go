//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertAuditLog — db.AuditLog row → *api.AuditLog 변환
//ff:checked llm=yongol-gen hash=42cd2033
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/ssac/pkg/pgtypex"
)

func convertAuditLog(row db.AuditLog) (*api.AuditLog, error) {
	return &api.AuditLog{
		Action: ptrOf(row.Action),
		ActorId: ptrOf(row.ActorID),
		CreatedAt: ptrOf(pgtypex.FromPgTimestamptz(row.CreatedAt)),
		Detail: ptrOf(row.Detail),
		Id: ptrOf(row.ID),
		OrgId: ptrOf(row.OrgID),
		ResourceId: ptrOf(row.ResourceID),
		ResourceType: ptrOf(row.ResourceType),
	}, nil
}
