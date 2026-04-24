//ff:func feature=service type=handler control=sequence
//ff:what ExecuteWorkflow — HTTP handler
//ff:checked llm=yongol-gen hash=ee83697f
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

func (server *Server) ExecuteWorkflow(ctx context.Context, request api.ExecuteWorkflowRequestObject) (api.ExecuteWorkflowResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ExecuteWorkflow")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ExecuteWorkflow")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ExecuteWorkflow")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "ExecuteWorkflow", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 404)
		return api.ExecuteWorkflow404JSONResponse{Error: "Workflow not found", Code: strPtr("not_found")}, nil
	}

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Tx: nil, Action: "ExecuteWorkflow", Resource: "workflow", Claim: currentUser, ResourceID: wf.ID})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 403, "err", err)
		return api.ExecuteWorkflow403JSONResponse{Error: "Forbidden", Code: strPtr("forbidden")}, nil
	}

	if !statemachine.WorkflowCanTransition(wf.Status, "ExecuteWorkflow") {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 409)
		return api.ExecuteWorkflow409JSONResponse{Error: "Cannot execute", Code: strPtr("conflict")}, nil
	}

	org, err := qtx.OrganizationFindByID(ctx, wf.OrgID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if org.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 404)
		return api.ExecuteWorkflow404JSONResponse{Error: "Organization not found", Code: strPtr("not_found")}, nil
	}

	sp, err := billing.Spend(billing.SpendRequest{Amount: 1, Current: org.CreditsBalance})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 402, "err", err)
		return api.ExecuteWorkflow402JSONResponse{Error: "Payment required", Code: strPtr("payment_required")}, nil
	}

	err = qtx.OrganizationUpdateCredits(ctx, db.OrganizationUpdateCreditsParams{CreditsBalance: sp.NewBalance, ID: org.ID})
	if err != nil { return nil, err }

	log, err := qtx.ExecutionLogCreate(ctx, db.ExecutionLogCreateParams{CreditsSpent: 1, OrgID: wf.OrgID, Status: "completed", WorkflowID: wf.ID})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	logConverted, err := convertExecutionLog(log)
	if err != nil { return nil, err }
	return api.ExecuteWorkflow200JSONResponse{
		Log: logConverted,
	}, nil
}
