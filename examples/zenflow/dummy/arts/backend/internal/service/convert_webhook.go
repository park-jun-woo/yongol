//ff:func feature=service type=util control=sequence topic=response-serialize
//ff:what convertWebhook — db.Webhook row → *api.Webhook 변환
//ff:checked llm=yongol-gen hash=442f3581
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/ssac/pkg/pgtypex"
)

func convertWebhook(row db.Webhook) (*api.Webhook, error) {
	return &api.Webhook{
		CreatedAt: ptrOf(pgtypex.FromPgTimestamptz(row.CreatedAt)),
		EventType: ptrOf(row.EventType),
		Id: ptrOf(row.ID),
		OrgId: ptrOf(row.OrgID),
		Url: ptrOf(row.Url),
	}, nil
}
