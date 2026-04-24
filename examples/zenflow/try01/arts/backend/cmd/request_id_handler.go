//ff:type feature=main type=util topic=request-id
//ff:what requestIDHandler — ctx 에 심긴 request_id 를 slog 레코드에 자동 주입
//ff:checked llm=yongol-gen hash=db4041fe


package main

import (
	"context"
	"log/slog"

	"github.com/park-jun-woo/zenflow/internal/middleware"
)

// requestIDHandler wraps a base slog.Handler and augments every record with
// the "request_id" attribute sourced from the context.Context. Records whose
// context carries no id are passed through unchanged.
type requestIDHandler struct {
	slog.Handler
}

// Handle implements slog.Handler. It adds request_id (when present) without
// mutating attrs the user supplied through slog.With() or log-site args.
func (h *requestIDHandler) Handle(ctx context.Context, r slog.Record) error {
	if id := middleware.RequestIDFromStdContext(ctx); id != "" {
		r.AddAttrs(slog.String("request_id", id))
	}
	return h.Handler.Handle(ctx, r)
}

// WithAttrs preserves the wrapper while delegating attribute cloning.
func (h *requestIDHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &requestIDHandler{Handler: h.Handler.WithAttrs(attrs)}
}

// WithGroup preserves the wrapper while delegating group scoping.
func (h *requestIDHandler) WithGroup(name string) slog.Handler {
	return &requestIDHandler{Handler: h.Handler.WithGroup(name)}
}
