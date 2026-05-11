//ff:func feature=service type=util control=iteration dimension=1 topic=response-serialize
//ff:what convertExecutionLogList — []db.ExecutionLog → []api.ExecutionLog 변환
//ff:checked llm=yongol-gen hash=fc33597b
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
)

func convertExecutionLogList(rows []db.ExecutionLog) ([]api.ExecutionLog, error) {
	result := make([]api.ExecutionLog, len(rows))
	for i, row := range rows {
		item, err := convertExecutionLog(row)
		if err != nil {
			return nil, err
		}
		result[i] = *item
	}
	return result, nil
}
