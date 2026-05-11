//ff:func feature=service type=handler control=sequence
//ff:what PauseWorkflow — Pause a workflow
//ff:checked llm=yongol-gen hash=79be217b
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
	"github.com/park-jun-woo/zenflow-try01/internal/statemachine"
	"log/slog"
	"strconv"
)

func (server *Server) PauseWorkflow(ctx context.Context, request api.PauseWorkflowRequestObject) (api.PauseWorkflowResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "PauseWorkflow")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "PauseWorkflow")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=PauseWorkflow")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "PauseWorkflow", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	ownerWorkflow, err := qtx.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "PauseWorkflow", "status", 403, "err", err)
		return api.PauseWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "PauseWorkflow", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "PauseWorkflow", "status", 403, "err", err)
		return api.PauseWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "PauseWorkflow", "status", 404)
		return api.PauseWorkflow404JSONResponse{Error: "Workflow not found", Code: "not_found"}, nil
	}

	if !statemachine.WorkflowCanTransition(wf.Status, "PauseWorkflow") {
		slog.Warn("handler: 4xx", "op", "PauseWorkflow", "status", 409)
		return api.PauseWorkflow409JSONResponse{Error: "Cannot pause workflow", Code: "conflict"}, nil
	}

	err = qtx.WorkflowUpdateStatus(ctx, db.WorkflowUpdateStatusParams{ID: wf.ID, Status: "paused"})
	if err != nil { return nil, err }

	updated, err := qtx.WorkflowFindByID(ctx, wf.ID)
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	workflowConverted, err := convertWorkflow(updated)
	if err != nil { return nil, err }
	return api.PauseWorkflow200JSONResponse{
		Workflow: workflowConverted,
	}, nil
}
