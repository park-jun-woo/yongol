//ff:func feature=service type=util control=iteration dimension=1 topic=response-serialize
//ff:what convertWorkflowList — []db.Workflow → []api.Workflow 변환
//ff:checked llm=yongol-gen hash=32e576eb
package service

import (
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
)

func convertWorkflowList(rows []db.Workflow) ([]api.Workflow, error) {
	result := make([]api.Workflow, len(rows))
	for i, row := range rows {
		item, err := convertWorkflow(row)
		if err != nil {
			return nil, err
		}
		result[i] = *item
	}
	return result, nil
}
