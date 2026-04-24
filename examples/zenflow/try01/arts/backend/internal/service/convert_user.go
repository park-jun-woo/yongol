//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertUser — db.User row → *api.User 변환
//ff:checked llm=yongol-gen hash=8d5864fc
package service

import (
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
)

func convertUser(row db.User) (*api.User, error) {
	return &api.User{
		CreatedAt: ptrOf(row.CreatedAt.Time),
		Email: ptrOf(row.Email),
		Id: ptrOf(row.ID),
		Name: ptrOf(row.Name),
		OrgId: ptrOf(row.OrgID),
		Role: ptrOf(row.Role),
	}, nil
}
