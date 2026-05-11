//ff:func feature=service type=util control=iteration dimension=1 topic=response-serialize
//ff:what convertAuditLogList — []db.AuditLog → []api.AuditLog 변환
//ff:checked llm=yongol-gen hash=369bba34
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
)

func convertAuditLogList(rows []db.AuditLog) ([]api.AuditLog, error) {
	result := make([]api.AuditLog, len(rows))
	for i, row := range rows {
		item, err := convertAuditLog(row)
		if err != nil {
			return nil, err
		}
		result[i] = *item
	}
	return result, nil
}
