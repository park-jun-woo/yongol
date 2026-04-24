//ff:func feature=service type=handler control=sequence
//ff:what CreateWorkflow — HTTP handler
//ff:checked llm=yongol-gen hash=ea6d864f
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

func (server *Server) CreateWorkflow(ctx context.Context, request api.CreateWorkflowRequestObject) (api.CreateWorkflowResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "CreateWorkflow")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "CreateWorkflow")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=CreateWorkflow")
	}

	tx, err := server.DB.BeginTx(ctx, nil)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			slog.Warn("rollback failed", "op", "CreateWorkflow", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Tx: tx, Action: "CreateWorkflow", Resource: "workflow", Claim: currentUser, })
	if err != nil {
		slog.Warn("handler: 4xx", "op", "CreateWorkflow", "status", 403, "err", err)
		return api.CreateWorkflow403JSONResponse{Error: "Forbidden", Code: strPtr("forbidden")}, nil
	}

	workflow, err := qtx.WorkflowCreate(ctx, db.WorkflowCreateParams{OrgID: currentUser.OrgID, Status: "draft", Title: request.Body.Title, TriggerEvent: request.Body.TriggerEvent})
	if err != nil { return nil, err }

	if err := tx.Commit(); err != nil { return nil, err }

	workflowConverted, err := convertWorkflow(workflow)
	if err != nil { return nil, err }
	return api.CreateWorkflow200JSONResponse{
		Workflow: workflowConverted,
	}, nil
}
