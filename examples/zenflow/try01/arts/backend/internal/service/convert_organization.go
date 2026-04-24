//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertOrganization — db.Organization row → *api.Organization 변환
//ff:checked llm=yongol-gen hash=54ae0714
package service

import (
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
)

func convertOrganization(row db.Organization) (*api.Organization, error) {
	return &api.Organization{
		CreatedAt: ptrOf(row.CreatedAt.Time),
		CreditsBalance: ptrOf(row.CreditsBalance),
		Id: ptrOf(row.ID),
		Name: ptrOf(row.Name),
		PlanType: ptrOf(row.PlanType),
	}, nil
}
