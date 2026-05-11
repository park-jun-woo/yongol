//ff:func feature=service type=handler control=sequence
//ff:what ListTemplates — List templates (cursor pagination)
//ff:checked llm=yongol-gen hash=6d4c3b5f
package service

import (
	"context"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"log/slog"
)

func (server *Server) ListTemplates(ctx context.Context, request api.ListTemplatesRequestObject) (api.ListTemplatesResponseObject, error) {
	slog.DebugContext(ctx, "handler entry", "op", "ListTemplates")

	items, err := server.Queries.TemplateListCursor(ctx, db.TemplateListCursorParams{Cursor: derefInt64(request.Params.Cursor), FilterCategory: derefStr(request.Params.Category), PerPage: derefInt32(request.Params.PerPage)})
	if err != nil { return nil, err }

	itemsConverted, err := convertTemplateList(items)
	if err != nil { return nil, err }
	return api.ListTemplates200JSONResponse{
		Items: itemsConverted,
	}, nil
}
