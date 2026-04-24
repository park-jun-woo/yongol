//ff:func feature=service type=handler control=sequence
//ff:what AddAction — HTTP handler
//ff:checked llm=yongol-gen hash=74354eb8
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
	"github.com/park-jun-woo/zenflow/internal/model"
	"log/slog"
)

func (server *Server) AddAction(ctx context.Context, request api.AddActionRequestObject) (api.AddActionResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "AddAction")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "AddAction")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=AddAction")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "AddAction", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "AddAction", "status", 404)
		return api.AddAction404JSONResponse{Error: "Workflow not found", Code: strPtr("not_found")}, nil
	}

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Tx: nil, Action: "AddAction", Resource: "workflow", Claim: currentUser, ResourceID: wf.ID})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "AddAction", "status", 403, "err", err)
		return api.AddAction403JSONResponse{Error: "Forbidden", Code: strPtr("forbidden")}, nil
	}

	action, err := qtx.ActionCreate(ctx, db.ActionCreateParams{ActionType: request.Body.ActionType, PayloadTemplate: request.Body.PayloadTemplate, SequenceOrder: request.Body.SequenceOrder, WorkflowID: wf.ID})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	actionConverted, err := convertAction(action)
	if err != nil { return nil, err }
	return api.AddAction200JSONResponse{
		Action: actionConverted,
	}, nil
}
