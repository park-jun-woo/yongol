//ff:func feature=service type=handler control=sequence
//ff:what GetDashboard — Get org dashboard summary
//ff:checked llm=yongol-gen hash=27015578
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
)

func (server *Server) GetDashboard(ctx context.Context, request api.GetDashboardRequestObject) (api.GetDashboardResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "GetDashboard")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "GetDashboard")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=GetDashboard")
	}

	_, err := authz.Check(authz.CheckRequest{Ctx: ctx, Action: "GetDashboard", Resource: "dashboard", Claim: currentUser, Owners: nil})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "GetDashboard", "status", 403, "err", err)
		return api.GetDashboard403JSONResponse{Error: "Forbidden", Code: "forbidden"}, nil
	}

	org, err := server.Queries.OrganizationFindByID(ctx, currentUser.OrgID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if org.ID == 0 {
		slog.Warn("handler: 4xx", "op", "GetDashboard", "status", 404)
		return api.GetDashboard404JSONResponse{Error: "Organization not found", Code: "not_found"}, nil
	}

	summaryConverted, err := convertOrganization(org)
	if err != nil { return nil, err }
	return api.GetDashboard200JSONResponse{
		Summary: summaryConverted,
	}, nil
}
