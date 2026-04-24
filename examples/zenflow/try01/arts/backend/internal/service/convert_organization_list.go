//ff:func feature=service type=util control=iteration dimension=1 topic=response-serialize
//ff:what convertOrganizationList — []db.Organization → []api.Organization 변환
//ff:checked llm=yongol-gen hash=585ff739
package service

import (
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
)

func convertOrganizationList(rows []db.Organization) ([]api.Organization, error) {
	result := make([]api.Organization, len(rows))
	for i, row := range rows {
		item, err := convertOrganization(row)
		if err != nil {
			return nil, err
		}
		result[i] = *item
	}
	return result, nil
}
