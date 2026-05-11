//ff:func feature=service type=handler control=sequence
//ff:what CreateAction — Create an action for a workflow
//ff:checked llm=yongol-gen hash=e4e60044
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
	"strconv"
)

func (server *Server) CreateAction(ctx context.Context, request api.CreateActionRequestObject) (api.CreateActionResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "CreateAction")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "CreateAction")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=CreateAction")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "CreateAction", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	ownerWorkflow, err := qtx.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "CreateAction", "status", 403, "err", err)
		return api.CreateAction403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "CreateAction", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "CreateAction", "status", 403, "err", err)
		return api.CreateAction403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "CreateAction", "status", 404)
		return api.CreateAction404JSONResponse{Error: "Workflow not found", Code: "not_found"}, nil
	}

	action, err := qtx.ActionCreate(ctx, db.ActionCreateParams{ActionType: request.Body.ActionType, Config: request.Body.Config, SequenceOrder: request.Body.SequenceOrder, WorkflowID: wf.ID})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	actionConverted, err := convertAction(action)
	if err != nil { return nil, err }
	return api.CreateAction201JSONResponse{
		Action: actionConverted,
	}, nil
}
