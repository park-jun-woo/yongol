//ff:func feature=service type=handler control=sequence
//ff:what ActivateWorkflow — Activate a workflow
//ff:checked llm=yongol-gen hash=a8ab7392
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/billing"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/zenflow-try01/internal/model"
	"github.com/park-jun-woo/zenflow-try01/internal/statemachine"
	"log/slog"
	"strconv"
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

	ownerWorkflow, err := qtx.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 403, "err", err)
		return api.ActivateWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ActivateWorkflow", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 403, "err", err)
		return api.ActivateWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 404)
		return api.ActivateWorkflow404JSONResponse{Error: "Workflow not found", Code: "not_found"}, nil
	}

	org, err := qtx.OrganizationFindByID(ctx, wf.OrgID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if org.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 404)
		return api.ActivateWorkflow404JSONResponse{Error: "Organization not found", Code: "not_found"}, nil
	}

	if billing.IsZeroBalance(billing.IsZeroBalanceRequest{Balance: org.CreditsBalance}) {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 402)
		return api.ActivateWorkflow402JSONResponse{Error: "Insufficient credits", Code: "payment_required"}, nil
	}

	if !statemachine.WorkflowCanTransition(wf.Status, "ActivateWorkflow") {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 409)
		return api.ActivateWorkflow409JSONResponse{Error: "Cannot activate workflow", Code: "conflict"}, nil
	}

	err = qtx.WorkflowUpdateStatus(ctx, db.WorkflowUpdateStatusParams{ID: wf.ID, Status: "active"})
	if err != nil { return nil, err }

	updated, err := qtx.WorkflowFindByID(ctx, wf.ID)
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	workflowConverted, err := convertWorkflow(updated)
	if err != nil { return nil, err }
	return api.ActivateWorkflow200JSONResponse{
		Workflow: workflowConverted,
	}, nil
}
