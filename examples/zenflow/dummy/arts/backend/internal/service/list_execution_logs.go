//ff:func feature=service type=handler control=sequence
//ff:what ListExecutionLogs — List execution logs for a workflow
//ff:checked llm=yongol-gen hash=50939db1
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

func (server *Server) ListExecutionLogs(ctx context.Context, request api.ListExecutionLogsRequestObject) (api.ListExecutionLogsResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ListExecutionLogs")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ListExecutionLogs")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ListExecutionLogs")
	}

	ownerWorkflow, err := server.Queries.OwnerLookupWorkflow(ctx, request.Id)
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListExecutionLogs", "status", 403, "err", err)
		return api.ListExecutionLogs403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}
	_, err = authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ListExecutionLogs", Resource: "workflow", Claim: currentUser, ResourceID: strconv.FormatInt(request.Id, 10), Owners: map[string]map[string]any{"workflow": {fmt.Sprint(request.Id): ownerWorkflow}}})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListExecutionLogs", "status", 403, "err", err)
		return api.ListExecutionLogs403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	logs, err := server.Queries.ExecutionLogListByWorkflowID(ctx, request.Id)
	if err != nil { return nil, err }

	executionLogsConverted, err := convertExecutionLogList(logs)
	if err != nil { return nil, err }
	return api.ListExecutionLogs200JSONResponse{
		ExecutionLogs: ptrOf(executionLogsConverted),
	}, nil
}
