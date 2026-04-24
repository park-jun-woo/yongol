//ff:func feature=service type=handler control=sequence
//ff:what ListWorkflows — HTTP handler
//ff:checked llm=yongol-gen hash=525995ba
package service

import (
	"context"
	"fmt"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/model"
	"log/slog"
)

func (server *Server) ListWorkflows(ctx context.Context, request api.ListWorkflowsRequestObject) (api.ListWorkflowsResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ListWorkflows")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ListWorkflows")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ListWorkflows")
	}

	_, err := authz.Check(authz.CheckRequest{Ctx: ctx, Tx: nil, Action: "ListWorkflows", Resource: "workflow", Claim: currentUser, })
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListWorkflows", "status", 403, "err", err)
		return api.ListWorkflows403JSONResponse{Error: "Forbidden", Code: strPtr("forbidden")}, nil
	}

	workflows, err := server.Queries.WorkflowListByOrgID(ctx, currentUser.OrgID)
	if err != nil { return nil, err }

	workflowsConverted, err := convertWorkflowList(workflows)
	if err != nil { return nil, err }
	return api.ListWorkflows200JSONResponse{
		Workflows: ptrOf(workflowsConverted),
	}, nil
}
