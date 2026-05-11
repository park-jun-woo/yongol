//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertUser — db.User row → *api.User 변환
//ff:checked llm=yongol-gen hash=79cb3964
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/ssac/pkg/pgtypex"
)

func convertUser(row db.User) (*api.User, error) {
	return &api.User{
		CreatedAt: ptrOf(pgtypex.FromPgTimestamptz(row.CreatedAt)),
		Email: ptrOf(row.Email),
		Id: ptrOf(row.ID),
		OrgId: ptrOf(row.OrgID),
		Role: ptrOf(row.Role),
	}, nil
}
