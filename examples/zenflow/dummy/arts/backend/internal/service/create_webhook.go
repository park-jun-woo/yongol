//ff:func feature=service type=handler control=sequence
//ff:what CreateWebhook — Register a webhook URL
//ff:checked llm=yongol-gen hash=7ce3468d
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/zenflow-try01/internal/model"
	"log/slog"
)

func (server *Server) CreateWebhook(ctx context.Context, request api.CreateWebhookRequestObject) (api.CreateWebhookResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "CreateWebhook")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "CreateWebhook")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=CreateWebhook")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "CreateWebhook", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "CreateWebhook", Resource: "webhook", Claim: currentUser, Owners: nil})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "CreateWebhook", "status", 403, "err", err)
		return api.CreateWebhook403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wh, err := qtx.WebhookCreate(ctx, db.WebhookCreateParams{EventType: request.Body.EventType, OrgID: currentUser.OrgID, Url: request.Body.Url})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	webhookConverted, err := convertWebhook(wh)
	if err != nil { return nil, err }
	return api.CreateWebhook201JSONResponse{
		Webhook: webhookConverted,
	}, nil
}
