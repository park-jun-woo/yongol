//ff:func feature=service type=handler control=sequence
//ff:what ListWorkflowVersions — List all versions of a workflow
//ff:checked llm=yongol-gen hash=061dca53
package service

import (
	"context"
	"fmt"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/zenflow-try01/internal/model"
	"log/slog"
	"strconv"
)

func (server *Server) ListWorkflowVersions(ctx context.Context, request api.ListWorkflowVersionsRequestObject) (api.ListWorkflowVersionsResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ListWorkflowVersions")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ListWorkflowVersions")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ListWorkflowVersions")
	}

	ownerWorkflow, err := server.Queries.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListWorkflowVersions", "status", 403, "err", err)
		return api.ListWorkflowVersions403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ListWorkflowVersions", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListWorkflowVersions", "status", 403, "err", err)
		return api.ListWorkflowVersions403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	workflows, err := server.Queries.WorkflowListVersions(ctx, db.WorkflowListVersionsParams{OrgID: currentUser.OrgID, RootID: request.Id})
	if err != nil { return nil, err }

	workflowsConverted, err := convertWorkflowList(workflows)
	if err != nil { return nil, err }
	return api.ListWorkflowVersions200JSONResponse{
		Workflows: ptrOf(workflowsConverted),
	}, nil
}
