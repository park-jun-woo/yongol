//ff:func feature=service type=util control=iteration dimension=1 topic=response-serialize
//ff:what convertActionList — []db.Action → []api.Action 변환
//ff:checked llm=yongol-gen hash=5716fb75
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
)

func convertActionList(rows []db.Action) ([]api.Action, error) {
	result := make([]api.Action, len(rows))
	for i, row := range rows {
		item, err := convertAction(row)
		if err != nil {
			return nil, err
		}
		result[i] = *item
	}
	return result, nil
}
