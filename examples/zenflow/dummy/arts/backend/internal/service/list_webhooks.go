//ff:func feature=service type=handler control=sequence
//ff:what ListWebhooks — List org webhooks
//ff:checked llm=yongol-gen hash=7c654ef5
package service

import (
	"context"
	"fmt"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/model"
	"log/slog"
)

func (server *Server) ListWebhooks(ctx context.Context, request api.ListWebhooksRequestObject) (api.ListWebhooksResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ListWebhooks")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ListWebhooks")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ListWebhooks")
	}

	_, err := authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ListWebhooks", Resource: "webhook", Claim: currentUser, Owners: nil})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListWebhooks", "status", 403, "err", err)
		return api.ListWebhooks403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	webhooks, err := server.Queries.WebhookListByOrgID(ctx, currentUser.OrgID)
	if err != nil { return nil, err }

	webhooksConverted, err := convertWebhookList(webhooks)
	if err != nil { return nil, err }
	return api.ListWebhooks200JSONResponse{
		Webhooks: ptrOf(webhooksConverted),
	}, nil
}
