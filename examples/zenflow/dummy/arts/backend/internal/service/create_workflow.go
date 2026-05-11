//ff:func feature=service type=handler control=sequence
//ff:what CreateWorkflow — Create a workflow
//ff:checked llm=yongol-gen hash=e8ae5474
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
)

func (server *Server) CreateWorkflow(ctx context.Context, request api.CreateWorkflowRequestObject) (api.CreateWorkflowResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "CreateWorkflow")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "CreateWorkflow")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=CreateWorkflow")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "CreateWorkflow", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "CreateWorkflow", Resource: "workflow", Claim: currentUser, Owners: nil})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "CreateWorkflow", "status", 403, "err", err)
		return api.CreateWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := qtx.WorkflowCreate(ctx, db.WorkflowCreateParams{OrgID: currentUser.OrgID, Title: request.Body.Title, TriggerEvent: request.Body.TriggerEvent})
	if err != nil { return nil, err }

	err = qtx.AuditLogCreate(ctx, db.AuditLogCreateParams{Action: "CreateWorkflow", ActorID: currentUser.ID, Detail: wf.Title, OrgID: currentUser.OrgID, ResourceID: wf.ID, ResourceType: "workflow"})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	workflowConverted, err := convertWorkflow(wf)
	if err != nil { return nil, err }
	return api.CreateWorkflow201JSONResponse{
		Workflow: workflowConverted,
	}, nil
}
