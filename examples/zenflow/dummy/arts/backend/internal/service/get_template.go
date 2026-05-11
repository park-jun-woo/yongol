//ff:func feature=service type=handler control=sequence
//ff:what GetTemplate — Get template detail
//ff:checked llm=yongol-gen hash=8ed43441
package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"log/slog"
)

func (server *Server) GetTemplate(ctx context.Context, request api.GetTemplateRequestObject) (api.GetTemplateResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "GetTemplate")

	tmpl, err := server.Queries.TemplateFindByID(ctx, request.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }

	if tmpl.ID == 0 {
		slog.Warn("handler: 4xx", "op", "GetTemplate", "status", 404)
		return api.GetTemplate404JSONResponse{Error: "Template not found", Code: "not_found"}, nil
	}

	templateConverted, err := convertTemplate(tmpl)
	if err != nil { return nil, err }
	return api.GetTemplate200JSONResponse{
		Template: templateConverted,
	}, nil
}
