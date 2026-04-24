//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertUser — db.User row → *api.User 변환
//ff:checked llm=yongol-gen hash=4c4a6847
package service

import (
	"github.com/example/zenflow_try01/internal/api"
	"github.com/example/zenflow_try01/internal/db"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func convertUser(row db.User) (*api.User, error) {
	return &api.User{
		CreatedAt: row.CreatedAt.Time,
		Email: openapi_types.Email(row.Email),
		Id: row.ID,
		Role: api.UserRole(row.Role),
	}, nil
}
