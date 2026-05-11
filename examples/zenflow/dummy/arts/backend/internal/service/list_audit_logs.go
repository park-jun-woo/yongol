//ff:func feature=service type=handler control=sequence
//ff:what ListAuditLogs — List audit logs with pagination
//ff:checked llm=yongol-gen hash=360dc0a5
package service

import (
	"context"
	"fmt"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/zenflow-try01/internal/model"
	"log/slog"
)

func (server *Server) ListAuditLogs(ctx context.Context, request api.ListAuditLogsRequestObject) (api.ListAuditLogsResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ListAuditLogs")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "ListAuditLogs")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=ListAuditLogs")
	}

	_, err := authz.Check(authz.CheckRequest{Ctx: ctx, Action: "ListAuditLogs", Resource: "audit_log", Claim: currentUser, Owners: nil})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "ListAuditLogs", "status", 403, "err", err)
		return api.ListAuditLogs403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	items, err := server.Queries.AuditLogListByOrgIDPaged(ctx, db.AuditLogListByOrgIDPagedParams{FilterAction: derefStr(request.Params.Action), OrgID: currentUser.OrgID, Page: derefInt32(request.Params.Page), PerPage: derefInt32(request.Params.PerPage), SortBy: derefEnum(request.Params.SortBy), SortDir: derefEnum(request.Params.SortDir)})
	if err != nil { return nil, err }

	total, err := server.Queries.AuditLogCountByOrgIDFiltered(ctx, db.AuditLogCountByOrgIDFilteredParams{FilterAction: derefStr(request.Params.Action), OrgID: currentUser.OrgID})
	if err != nil { return nil, err }

	itemsConverted, err := convertAuditLogList(items)
	if err != nil { return nil, err }
	return api.ListAuditLogs200JSONResponse{
		Items: itemsConverted,
		Total: total,
	}, nil
}
