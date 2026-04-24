//ff:func feature=service type=handler control=sequence
//ff:what ListWorkflows — HTTP handler
//ff:checked llm=yongol-gen hash=f5719dd8
package service

import (
	"context"
	"fmt"
	"github.com/example/zenflow_try01/internal/api"
	"github.com/example/zenflow_try01/internal/db"
	"github.com/example/zenflow_try01/internal/model"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"log/slog"
)

func (server *Server) ListWorkflows(ctx context.Context, request api.ListWorkflowsRequestObject) (api.ListWorkflowsResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ListWorkflows")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ListWorkflows")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ListWorkflows")
	}

	_, err := authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ListWorkflows", Resource: "workflow", Claim: currentUser, Owners: nil})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListWorkflows", "status", 403, "err", err)
		return api.ListWorkflows403JSONResponse{Error: "Forbidden", Code: strPtr("forbidden")}, nil
	}

	items, err := server.Queries.WorkflowListByOwnerID(ctx, db.WorkflowListByOwnerIDParams{OwnerID: currentUser.ID, Page: derefInt(request.Params.Page), PerPage: derefInt(request.Params.PerPage)})
	if err != nil { return nil, err }

	total, err := server.Queries.WorkflowCountByOwnerID(ctx, currentUser.ID)
	if err != nil { return nil, err }

	itemsConverted, err := convertWorkflowList(items)
	if err != nil { return nil, err }
	return api.ListWorkflows200JSONResponse{
		Items: itemsConverted,
		Total: total,
	}, nil
}
