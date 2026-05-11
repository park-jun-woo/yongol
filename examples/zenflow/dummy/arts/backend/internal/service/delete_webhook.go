//ff:func feature=service type=handler control=sequence
//ff:what DeleteWebhook — Delete a webhook
//ff:checked llm=yongol-gen hash=3409283c
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/model"
	"log/slog"
	"strconv"
)

func (server *Server) DeleteWebhook(ctx context.Context, request api.DeleteWebhookRequestObject) (api.DeleteWebhookResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "DeleteWebhook")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "DeleteWebhook")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=DeleteWebhook")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "DeleteWebhook", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	ownerWebhook, err := qtx.OwnerLookupWebhook(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "DeleteWebhook", "status", 403, "err", err)
		return api.DeleteWebhook403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "DeleteWebhook", Resource: "webhook", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"webhook": {fmt.Sprint(request.Id): ownerWebhook}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "DeleteWebhook", "status", 403, "err", err)
		return api.DeleteWebhook403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wh, err := qtx.WebhookFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wh.ID == 0 {
		slog.Warn("handler: 4xx", "op", "DeleteWebhook", "status", 404)
		return api.DeleteWebhook404JSONResponse{Error: "Webhook not found", Code: "not_found"}, nil
	}

	err = qtx.WebhookDelete(ctx, wh.ID)
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	return api.DeleteWebhook200JSONResponse{
		Deleted: ptrOf(true),
	}, nil
}
