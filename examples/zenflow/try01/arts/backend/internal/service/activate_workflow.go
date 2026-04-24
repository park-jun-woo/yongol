//ff:func feature=service type=handler control=sequence
//ff:what ActivateWorkflow — HTTP handler
//ff:checked llm=yongol-gen hash=2df0a04a
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/billing"
	"github.com/park-jun-woo/zenflow/internal/db"
	"github.com/park-jun-woo/zenflow/internal/model"
	"github.com/park-jun-woo/zenflow/internal/statemachine"
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

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 404)
		return api.ActivateWorkflow404JSONResponse{Error: "Workflow not found", Code: strPtr("not_found")}, nil
	}

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Tx: nil, Action: "ActivateWorkflow", Resource: "workflow", Claim: currentUser, ResourceID: wf.ID})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 403, "err", err)
		return api.ActivateWorkflow403JSONResponse{Error: "Forbidden", Code: strPtr("forbidden")}, nil
	}

	org, err := qtx.OrganizationFindByID(ctx, wf.OrgID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if org.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 404)
		return api.ActivateWorkflow404JSONResponse{Error: "Organization not found", Code: strPtr("not_found")}, nil
	}

	_, err = billing.CheckCredits(billing.CheckCreditsRequest{Current: org.CreditsBalance})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 402, "err", err)
		return api.ActivateWorkflow402JSONResponse{Error: "Payment required", Code: strPtr("payment_required")}, nil
	}

	if !statemachine.WorkflowCanTransition(wf.Status, "ActivateWorkflow") {
		slog.Warn("handler: 4xx", "op", "ActivateWorkflow", "status", 409)
		return api.ActivateWorkflow409JSONResponse{Error: "Cannot activate", Code: strPtr("conflict")}, nil
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
