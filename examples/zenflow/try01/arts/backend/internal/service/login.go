//ff:func feature=service type=handler control=sequence
//ff:what Login — HTTP handler
//ff:checked llm=yongol-gen hash=89e602b6
package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/auth"
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/model"
	"log/slog"
)

func (server *Server) Login(ctx context.Context, request api.LoginRequestObject) (api.LoginResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "Login")

	user, err := server.Queries.UserFindByEmail(ctx, string(request.Body.Email))
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }
	if user.ID == 0 {
		_, _ = auth.VerifyPassword(auth.VerifyPasswordRequest{
			Password:     request.Body.Password,
			PasswordHash: auth.DummyHash,
		})
		slog.Warn("handler: 4xx", "op", "Login", "status", 401, "reason", "user not found")
		return api.Login401JSONResponse{Error: "Invalid credentials", Code: strPtr("unauthorized")}, nil
	}
	_, err = auth.VerifyPassword(auth.VerifyPasswordRequest{
		Password:     request.Body.Password,
		PasswordHash: user.PasswordHash,
	})
	if err != nil {
		slog.Warn("handler: 4xx", "op", "Login", "status", 401, "err", err)
		return api.Login401JSONResponse{Error: "Invalid credentials", Code: strPtr("unauthorized")}, nil
	}

	token, err := auth.IssueToken(auth.IssueTokenRequest{Claims: model.UserClaim{Email: user.Email, ID: user.ID, OrgID: user.OrgID, Role: user.Role}})
	if err != nil {
		slog.Error("handler: 5xx", "op", "Login", "status", 500, "err", err)
		return api.Login500JSONResponse{Error: "Internal error", Code: strPtr("internal_error")}, nil
	}

	return api.Login200JSONResponse{
		AccessToken: &token.AccessToken,
	}, nil
}
