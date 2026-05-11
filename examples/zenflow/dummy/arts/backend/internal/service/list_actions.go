//ff:func feature=service type=handler control=sequence
//ff:what ListActions — List actions for a workflow
//ff:checked llm=yongol-gen hash=a5e5797a
package service

import (
	"context"
	"fmt"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/model"
	"log/slog"
	"strconv"
)

func (server *Server) ListActions(ctx context.Context, request api.ListActionsRequestObject) (api.ListActionsResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ListActions")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ListActions")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ListActions")
	}

	ownerWorkflow, err := server.Queries.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListActions", "status", 403, "err", err)
		return api.ListActions403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ListActions", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListActions", "status", 403, "err", err)
		return api.ListActions403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	actions, err := server.Queries.ActionListByWorkflowID(ctx, request.Id)
	if err != nil { return nil, err }

	actionsConverted, err := convertActionList(actions)
	if err != nil { return nil, err }
	return api.ListActions200JSONResponse{
		Actions: ptrOf(actionsConverted),
	}, nil
}
