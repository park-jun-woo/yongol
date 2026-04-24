//ff:func feature=service type=handler control=sequence
//ff:what ActivateWorkflow — HTTP handler
//ff:checked llm=yongol-gen hash=8ad91274
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/example/zenflow_try01/internal/api"
	"github.com/example/zenflow_try01/internal/db"
	"github.com/example/zenflow_try01/internal/model"
	"github.com/example/zenflow_try01/internal/statemachine"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"log/slog"
)

func (server *Server) ActivateWorkflow(ctx context.Context, request api.ActivateWorkflowRequestObject) (api.ActivateWorkflowResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ActivateWorkflow")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ActivateWorkflow")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ActivateWorkflow")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "ActivateWorkflow", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ActivateWorkflow", Resource: "workflow", Claim: currentUser, ResourceID: request.Id, Owners: nil})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 403, "err", err)
		return api.ActivateWorkflow403JSONResponse{Error: "Forbidden", Code: strPtr("forbidden")}, nil
	}

	workflow, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if workflow.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 404)
		return api.ActivateWorkflow404JSONResponse{Error: "Workflow not found", Code: strPtr("not_found")}, nil
	}

	if !statemachine.WorkflowCanTransition(workflow.Status, "ActivateWorkflow") {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 409)
		return api.ActivateWorkflow409JSONResponse{Error: "Invalid state transition", Code: strPtr("conflict")}, nil
	}

	err = qtx.WorkflowUpdateStatus(ctx, db.WorkflowUpdateStatusParams{ID: workflow.ID, Status: "active"})
	if err != nil { return nil, err }

	updated, err := qtx.WorkflowFindByID(ctx, workflow.ID)
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	workflowConverted, err := convertWorkflow(updated)
	if err != nil { return nil, err }
	return api.ActivateWorkflow200JSONResponse{
		Workflow: *workflowConverted,
	}, nil
}
