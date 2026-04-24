//ff:func feature=service type=handler control=sequence
//ff:what Signup — HTTP handler
//ff:checked llm=yongol-gen hash=510339d7
package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/auth"
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
	"github.com/park-jun-woo/zenflow/internal/model"
	"log/slog"
)

func (server *Server) Signup(ctx context.Context, request api.SignupRequestObject) (api.SignupResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "Signup")

	tx, err := server.DB.Begin(ctx)
	if err != nil { return nil, err }
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			slog.Warn("rollback failed", "op", "Signup", "err", err)
		}
	}()
	qtx := server.Queries.WithTx(tx)

	existing, err := qtx.UserFindByEmail(ctx, string(request.Body.Email))
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if existing.ID != 0 {
		slog.Warn("handler: 4xx", "op", "Signup", "status", 409)
		return api.Signup409JSONResponse{Error: "Email already registered", Code: strPtr("conflict")}, nil
	}

	hp, err := auth.HashPassword(auth.HashPasswordRequest{Password: request.Body.Password})
	if err != nil {
		slog.Error("handler: 5xx", "op", "Signup", "status", 500, "err", err)
		return api.Signup500JSONResponse{Error: "Internal error", Code: strPtr("internal_error")}, nil
	}

	org, err := qtx.OrganizationCreate(ctx, db.OrganizationCreateParams{CreditsBalance: request.Body.CreditsBalance, Name: request.Body.OrgName, PlanType: string(request.Body.PlanType)})
	if err != nil { return nil, err }

	user, err := qtx.UserCreate(ctx, db.UserCreateParams{Email: string(request.Body.Email), Name: request.Body.Name, OrgID: org.ID, PasswordHash: hp.HashedPassword, Role: "admin"})
	if err != nil { return nil, err }

	token, err := auth.IssueToken(auth.IssueTokenRequest{Claims: model.UserClaim{Email: user.Email, ID: user.ID, OrgID: user.OrgID, Role: user.Role}})
	if err != nil {
		slog.Error("handler: 5xx", "op", "Signup", "status", 500, "err", err)
		return api.Signup500JSONResponse{Error: "Internal error", Code: strPtr("internal_error")}, nil
	}

	if err := tx.Commit(ctx); err != nil { return nil, err }

	organizationConverted, err := convertOrganization(org)
	if err != nil { return nil, err }
	userConverted, err := convertUser(user)
	if err != nil { return nil, err }
	return api.Signup200JSONResponse{
		AccessToken: &token.AccessToken,
		Organization: organizationConverted,
		User: userConverted,
	}, nil
}
