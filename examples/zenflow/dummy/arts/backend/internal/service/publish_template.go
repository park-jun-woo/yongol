//ff:func feature=service type=handler control=sequence
//ff:what PublishTemplate — Publish a workflow as a template
//ff:checked llm=yongol-gen hash=dfbef7b4
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
	"strconv"
)

func (server *Server) PublishTemplate(ctx context.Context, request api.PublishTemplateRequestObject) (api.PublishTemplateResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "PublishTemplate")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "PublishTemplate")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=PublishTemplate")
	}

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "PublishTemplate", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	ownerWorkflow, err := qtx.OwnerLookupWorkflow(ctx, request.Body.SourceWorkflowId)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "PublishTemplate", "status", 403, "err", err)
		return api.PublishTemplate403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "PublishTemplate", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Body.SourceWorkflowId, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Body.SourceWorkflowId): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "PublishTemplate", "status", 403, "err", err)
		return api.PublishTemplate403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := qtx.WorkflowFindByID(ctx, request.Body.SourceWorkflowId)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "PublishTemplate", "status", 404)
		return api.PublishTemplate404JSONResponse{Error: "Workflow not found", Code: "not_found"}, nil
	}

	existing, err := qtx.TemplateFindBySourceWorkflowID(ctx, wf.ID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if existing.ID != 0 {
		slog.Warn("handler: 4xx", "op", "PublishTemplate", "status", 409)
		return api.PublishTemplate409JSONResponse{Error: "Already published", Code: "conflict"}, nil
	}

	tmpl, err := qtx.TemplateCreate(ctx, db.TemplateCreateParams{Category: request.Body.Category, Description: request.Body.Description, OrgID: currentUser.OrgID, SourceWorkflowID: wf.ID, Title: request.Body.Title})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	templateConverted, err := convertTemplate(tmpl)
	if err != nil { return nil, err }
	return api.PublishTemplate201JSONResponse{
		Template: templateConverted,
	}, nil
}
