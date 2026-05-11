//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertTemplate — db.Template row → *api.Template 변환
//ff:checked llm=yongol-gen hash=2b400ded
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/ssac/pkg/pgtypex"
)

func convertTemplate(row db.Template) (*api.Template, error) {
	return &api.Template{
		Category: ptrOf(row.Category),
		CloneCount: ptrOf(row.CloneCount),
		CreatedAt: ptrOf(pgtypex.FromPgTimestamptz(row.CreatedAt)),
		Description: ptrOf(row.Description),
		Id: ptrOf(row.ID),
		OrgId: ptrOf(row.OrgID),
		SourceWorkflowId: ptrOf(row.SourceWorkflowID),
		Title: ptrOf(row.Title),
	}, nil
}
