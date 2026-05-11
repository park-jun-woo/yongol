//ff:func feature=service type=handler control=sequence
//ff:what CreateWorkflowVersion — Clone workflow into a new draft version
//ff:checked llm=yongol-gen hash=341a71a7
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
	"github.com/park-jun-woo/zenflow-try01/internal/versioning"
	"log/slog"
	"strconv"
)

func (server *Server) CreateWorkflowVersion(ctx context.Context, request api.CreateWorkflowVersionRequestObject) (api.CreateWorkflowVersionResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "CreateWorkflowVersion")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "CreateWorkflowVersion")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=CreateWorkflowVersion")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "CreateWorkflowVersion", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	ownerWorkflow, err := qtx.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "CreateWorkflowVersion", "status", 403, "err", err)
		return api.CreateWorkflowVersion403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "CreateWorkflowVersion", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "CreateWorkflowVersion", "status", 403, "err", err)
		return api.CreateWorkflowVersion403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := qtx.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "CreateWorkflowVersion", "status", 404)
		return api.CreateWorkflowVersion404JSONResponse{Error: "Workflow not found", Code: "not_found"}, nil
	}

	root, err := versioning.ResolveRootID(versioning.ResolveRootIDRequest{RootWorkflowID: wf.RootWorkflowID, WorkflowID: wf.ID})
	if err != nil {
		slog.Error("handler: 5xx", "op", "CreateWorkflowVersion", "status", 500, "err", err)
		return api.CreateWorkflowVersion500JSONResponse{Error: "Internal error", Code: "internal_error"}, nil
	}

	ver, err := versioning.NextVersion(versioning.NextVersionRequest{CurrentVersion: wf.Version})
	if err != nil {
		slog.Error("handler: 5xx", "op", "CreateWorkflowVersion", "status", 500, "err", err)
		return api.CreateWorkflowVersion500JSONResponse{Error: "Internal error", Code: "internal_error"}, nil
	}

	newWf, err := qtx.WorkflowCreateVersion(ctx, db.WorkflowCreateVersionParams{OrgID: wf.OrgID, RootWorkflowID: root.RootID, Title: wf.Title, TriggerEvent: wf.TriggerEvent, Version: ver.Version})
	if err != nil { return nil, err }

	err = qtx.ActionCopyToWorkflow(ctx, db.ActionCopyToWorkflowParams{SourceWorkflowID: wf.ID, TargetWorkflowID: newWf.ID})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	workflowConverted, err := convertWorkflow(newWf)
	if err != nil { return nil, err }
	return api.CreateWorkflowVersion201JSONResponse{
		Workflow: workflowConverted,
	}, nil
}
