//ff:func feature=service type=handler control=sequence
//ff:what Register — HTTP handler
//ff:checked llm=yongol-gen hash=ef4da53c
package service

import (
	"context"
	"errors"
	"github.com/example/zenflow_try01/internal/api"
	"github.com/example/zenflow_try01/internal/db"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/auth"
	"log/slog"
)

func (server *Server) Register(ctx context.Context, request api.RegisterRequestObject) (api.RegisterResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "Register")

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "Register", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	hp, err := auth.HashPassword(auth.HashPasswordRequest{Password: request.Body.Password})
	if err != nil {
		slog.Error("handler: 5xx", "op", "Register", "status", 500, "err", err)
		return api.Register500JSONResponse{Error: "Internal error", Code: strPtr("internal_error")}, nil
	}

	user, err := qtx.UserCreate(ctx, db.UserCreateParams{Email: string(request.Body.Email), PasswordHash: hp.HashedPassword, Role: "admin"})
	if err != nil { return nil, err }

	if err := tx.Commit(ctx); err != nil { return nil, err }

	userConverted, err := convertUser(user)
	if err != nil { return nil, err }
	return api.Register201JSONResponse{
		User: *userConverted,
	}, nil
}
