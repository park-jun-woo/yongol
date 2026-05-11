//ff:func feature=service type=handler control=sequence
//ff:what ExecuteWorkflow — Execute a workflow
//ff:checked llm=yongol-gen hash=d2a495e9
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
	"github.com/park-jun-woo/zenflow-try01/internal/worker"
	"log/slog"
	"strconv"
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

	ownerWorkflow, err := qtx.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 403, "err", err)
		return api.ExecuteWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ExecuteWorkflow", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 403, "err", err)
		return api.ExecuteWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 404)
		return api.ExecuteWorkflow404JSONResponse{Error: "Workflow not found", Code: "not_found"}, nil
	}

	if !statemachine.WorkflowCanTransition(wf.Status, "ExecuteWorkflow") {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 409)
		return api.ExecuteWorkflow409JSONResponse{Error: "Workflow is not active", Code: "conflict"}, nil
	}

	org, err := qtx.OrganizationFindByID(ctx, wf.OrgID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if org.ID == 0 {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 404)
		return api.ExecuteWorkflow404JSONResponse{Error: "Organization not found", Code: "not_found"}, nil
	}

	if billing.IsZeroBalance(billing.IsZeroBalanceRequest{Balance: org.CreditsBalance}) {
		slog.Warn("handler: 4xx", "op", "ExecuteWorkflow", "status", 402)
		return api.ExecuteWorkflow402JSONResponse{Error: "Insufficient credits", Code: "payment_required"}, nil
	}

	result, err := worker.ProcessActions(worker.ProcessActionsRequest{ActionCount: 0, WorkflowID: wf.ID})
	if err != nil {
		slog.Error("handler: 5xx", "op", "ExecuteWorkflow", "status", 500, "err", err)
		return api.ExecuteWorkflow500JSONResponse{Error: "Internal error", Code: "internal_error"}, nil
	}

	err = qtx.OrganizationDeductCredits(ctx, db.OrganizationDeductCreditsParams{Amount: 1, ID: wf.OrgID})
	if err != nil { return nil, err }

	log, err := qtx.ExecutionLogCreate(ctx, db.ExecutionLogCreateParams{CreditsSpent: 1, OrgID: wf.OrgID, Status: result.Status, WorkflowID: wf.ID})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	executionLogConverted, err := convertExecutionLog(log)
	if err != nil { return nil, err }
	return api.ExecuteWorkflow200JSONResponse{
		ExecutionLog: executionLogConverted,
	}, nil
}
