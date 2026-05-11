//ff:func feature=service type=handler control=sequence
//ff:what CloneTemplate — Clone a template into caller's org
//ff:checked llm=yongol-gen hash=7c8c25ec
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

func (server *Server) CloneTemplate(ctx context.Context, request api.CloneTemplateRequestObject) (api.CloneTemplateResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "CloneTemplate")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "CloneTemplate")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=CloneTemplate")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "CloneTemplate", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "CloneTemplate", Resource: "template", Claim: currentUser, Owners: nil})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "CloneTemplate", "status", 403, "err", err)
		return api.CloneTemplate403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	tmpl, err := qtx.TemplateFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if tmpl.ID == 0 {
		slog.Warn("handler: 4xx", "op", "CloneTemplate", "status", 404)
		return api.CloneTemplate404JSONResponse{Error: "Template not found", Code: "not_found"}, nil
	}

	sourceWf, err := qtx.WorkflowFindByID(ctx, tmpl.SourceWorkflowID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if sourceWf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "CloneTemplate", "status", 404)
		return api.CloneTemplate404JSONResponse{Error: "Source workflow not found", Code: "not_found"}, nil
	}

	newWf, err := qtx.WorkflowCreate(ctx, db.WorkflowCreateParams{OrgID: currentUser.OrgID, Title: tmpl.Title, TriggerEvent: sourceWf.TriggerEvent})
	if err != nil { return nil, err }

	err = qtx.ActionCopyToWorkflow(ctx, db.ActionCopyToWorkflowParams{SourceWorkflowID: sourceWf.ID, TargetWorkflowID: newWf.ID})
	if err != nil { return nil, err }

	err = qtx.TemplateIncrementCloneCount(ctx, tmpl.ID)
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	workflowConverted, err := convertWorkflow(newWf)
	if err != nil { return nil, err }
	return api.CloneTemplate201JSONResponse{
		Workflow: workflowConverted,
	}, nil
}
