//ff:func feature=service type=handler control=sequence
//ff:what ArchiveWorkflow — Archive a workflow
//ff:checked llm=yongol-gen hash=2b5c9337
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

func (server *Server) ArchiveWorkflow(ctx context.Context, request api.ArchiveWorkflowRequestObject) (api.ArchiveWorkflowResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ArchiveWorkflow")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ArchiveWorkflow")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ArchiveWorkflow")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "ArchiveWorkflow", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	ownerWorkflow, err := qtx.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ArchiveWorkflow", "status", 403, "err", err)
		return api.ArchiveWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ArchiveWorkflow", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ArchiveWorkflow", "status", 403, "err", err)
		return api.ArchiveWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ArchiveWorkflow", "status", 404)
		return api.ArchiveWorkflow404JSONResponse{Error: "Workflow not found", Code: "not_found"}, nil
	}

	if !statemachine.WorkflowCanTransition(wf.Status, "ArchiveWorkflow") {
		slog.Warn("handler: 4xx", "op", "ArchiveWorkflow", "status", 409)
		return api.ArchiveWorkflow409JSONResponse{Error: "Cannot archive workflow", Code: "conflict"}, nil
	}

	err = qtx.WorkflowUpdateStatus(ctx, db.WorkflowUpdateStatusParams{ID: wf.ID, Status: "archived"})
	if err != nil { return nil, err }

	updated, err := qtx.WorkflowFindByID(ctx, wf.ID)
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	workflowConverted, err := convertWorkflow(updated)
	if err != nil { return nil, err }
	return api.ArchiveWorkflow200JSONResponse{
		Workflow: workflowConverted,
	}, nil
}
