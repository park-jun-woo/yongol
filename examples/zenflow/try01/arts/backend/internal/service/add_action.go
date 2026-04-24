//ff:func feature=service type=handler control=sequence
//ff:what AddAction — HTTP handler
//ff:checked llm=yongol-gen hash=58b1cbd1
package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	tx, err := server.DB.BeginTx(ctx, nil)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			slog.Warn("rollback failed", "op", "AddAction", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "AddAction", "status", 404)
		return api.AddAction404JSONResponse{Error: "Workflow not found", Code: strPtr("not_found")}, nil
	}

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Tx: tx, Action: "AddAction", Resource: "workflow", Claim: currentUser, ResourceID: wf.ID})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "AddAction", "status", 403, "err", err)
		return api.AddAction403JSONResponse{Error: "Forbidden", Code: strPtr("forbidden")}, nil
	}

	action, err := qtx.ActionCreate(ctx, db.ActionCreateParams{ActionType: request.Body.ActionType, PayloadTemplate: request.Body.PayloadTemplate, SequenceOrder: request.Body.SequenceOrder, WorkflowID: wf.ID})
	if err != nil { return nil, err }

	if err := tx.Commit(); err != nil { return nil, err }

	actionConverted, err := convertAction(action)
	if err != nil { return nil, err }
	return api.AddAction200JSONResponse{
		Action: actionConverted,
	}, nil
}
