//ff:func feature=gen-gogin type=generator control=sequence topic=request-id
//ff:what loggerHelperRequestIDHandlerFile — cmd/request_id_handler.go 독립 파일 소스 반환

package boot

import (
	"fmt"
	"os"
	"path/filepath"
)

// writeRequestIDHandlerFile emits cmd/request_id_handler.go — a slog.Handler
// wrapper that reads middleware.RequestIDFromStdContext(ctx) and injects the
// correlation id as a slog attribute on every record. Emitted as a dedicated
// file (not through writeEnvHelperFiles) because it introduces a type +
// multiple methods rather than a single top-level func, which the per-func
// helper pipeline does not handle.
//
// modulePath is the generated project's go module so the import path for
// internal/middleware is correct.
func writeRequestIDHandlerFile(artifactsDir, modulePath string) error {
	if modulePath == "" {
		return nil
	}
	outDir := filepath.Join(artifactsDir, "backend", "cmd")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	src := fmt.Sprintf(`%s%s%s

package main

import (
	"context"
	"log/slog"

	"%s/internal/middleware"
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
`,
		"//"+"ff:type feature=main type=util topic=request-id\n",
		"//"+"ff:what requestIDHandler — ctx 에 심긴 request_id 를 slog 레코드에 자동 주입\n",
		"",
		modulePath)
	return os.WriteFile(filepath.Join(outDir, "request_id_handler.go"), []byte(src), 0o644)
}
