//ff:func feature=service type=util control=iteration dimension=1 topic=response-serialize
//ff:what convertWebhookList — []db.Webhook → []api.Webhook 변환
//ff:checked llm=yongol-gen hash=d81ae086
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
)

func convertWebhookList(rows []db.Webhook) ([]api.Webhook, error) {
	result := make([]api.Webhook, len(rows))
	for i, row := range rows {
		item, err := convertWebhook(row)
		if err != nil {
			return nil, err
		}
		result[i] = *item
	}
	return result, nil
}
