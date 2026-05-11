//ff:func feature=service type=handler control=sequence
//ff:what GetWorkflow — Get a workflow by ID
//ff:checked llm=yongol-gen hash=896c24f0
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/model"
	"log/slog"
	"strconv"
)

func (server *Server) GetWorkflow(ctx context.Context, request api.GetWorkflowRequestObject) (api.GetWorkflowResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "GetWorkflow")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "GetWorkflow")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=GetWorkflow")
	}

	ownerWorkflow, err := server.Queries.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "GetWorkflow", "status", 403, "err", err)
		return api.GetWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "GetWorkflow", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "GetWorkflow", "status", 403, "err", err)
		return api.GetWorkflow403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	wf, err := server.Queries.WorkflowFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if wf.ID == 0 {
		slog.Warn("handler: 4xx", "op", "GetWorkflow", "status", 404)
		return api.GetWorkflow404JSONResponse{Error: "Workflow not found", Code: "not_found"}, nil
	}

	workflowConverted, err := convertWorkflow(wf)
	if err != nil { return nil, err }
	return api.GetWorkflow200JSONResponse{
		Workflow: workflowConverted,
	}, nil
}
