//ff:func feature=service type=handler control=sequence
//ff:what GetCurrentUser — HTTP handler
//ff:checked llm=yongol-gen hash=47dd05dc
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/model"
	"log/slog"
)

func (server *Server) GetCurrentUser(ctx context.Context, request api.GetCurrentUserRequestObject) (api.GetCurrentUserResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "GetCurrentUser")
	currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)
	if !ok || currentUser == nil {
		slog.Error("missing currentUser in authenticated handler", "op", "GetCurrentUser")
		return nil, fmt.Errorf("missing currentUser in authenticated handler: op=GetCurrentUser")
	}

	user, err := server.Queries.UserFindByID(ctx, currentUser.ID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if user.ID == 0 {
		slog.Warn("handler: 4xx", "op", "GetCurrentUser", "status", 404)
		return api.GetCurrentUser404JSONResponse{Error: "User not found", Code: strPtr("not_found")}, nil
	}

	userConverted, err := convertUser(user)
	if err != nil { return nil, err }
	return api.GetCurrentUser200JSONResponse{
		User: userConverted,
	}, nil
}
