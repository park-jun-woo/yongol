//ff:func feature=gen-gogin type=generator control=iteration dimension=1 topic=request-id
//ff:what writeRequestIDHandlerFiles — cmd/request_id_handler*.go 4파일 emit (1 file 1 method)

package boot

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/ffhash"
)

// writeRequestIDHandlerFile emits cmd/request_id_handler*.go — a slog.Handler
// wrapper split into 1 type file + 3 method files (filefunc F3).
func writeRequestIDHandlerFile(artifactsDir, modulePath string) error {
	if modulePath == "" {
		return nil
	}
	outDir := filepath.Join(artifactsDir, "backend", "cmd")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}

	typeFile := fmt.Sprintf(`%s%s
package main

import "log/slog"

// requestIDHandler wraps a base slog.Handler and augments every record with
// the "request_id" attribute sourced from the context.Context. Records whose
// context carries no id are passed through unchanged.
type requestIDHandler struct {
	slog.Handler
}
`,
		"//"+"ff:type feature=main type=util topic=request-id\n",
		"//"+"ff:what requestIDHandler — ctx 에 심긴 request_id 를 slog 레코드에 자동 주입\n")

	handleFile := fmt.Sprintf(`%s%s
package main

import (
	"context"
	"log/slog"

	"%s/internal/middleware"
)

// Handle implements slog.Handler. It adds request_id (when present) without
// mutating attrs the user supplied through slog.With() or log-site args.
func (h *requestIDHandler) Handle(ctx context.Context, r slog.Record) error {
	if id := middleware.RequestIDFromStdContext(ctx); id != "" {
		r.AddAttrs(slog.String("request_id", id))
	}
	return h.Handler.Handle(ctx, r)
}
`,
		"//"+"ff:func feature=main type=util control=sequence topic=request-id\n",
		"//"+"ff:what requestIDHandler.Handle — request_id 를 slog Record 에 주입\n",
		modulePath)

	withAttrsFile := fmt.Sprintf(`%s%s
package main

import "log/slog"

// WithAttrs preserves the wrapper while delegating attribute cloning.
func (h *requestIDHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &requestIDHandler{Handler: h.Handler.WithAttrs(attrs)}
}
`,
		"//"+"ff:func feature=main type=util control=sequence topic=request-id\n",
		"//"+"ff:what requestIDHandler.WithAttrs — wrapper 보존하며 attribute 위임\n")

	withGroupFile := fmt.Sprintf(`%s%s
package main

import "log/slog"

// WithGroup preserves the wrapper while delegating group scoping.
func (h *requestIDHandler) WithGroup(name string) slog.Handler {
	return &requestIDHandler{Handler: h.Handler.WithGroup(name)}
}
`,
		"//"+"ff:func feature=main type=util control=sequence topic=request-id\n",
		"//"+"ff:what requestIDHandler.WithGroup — wrapper 보존하며 group 스코핑 위임\n")

	files := map[string]string{
		"request_id_handler.go":            typeFile,
		"request_id_handler_handle.go":     handleFile,
		"request_id_handler_with_attrs.go": withAttrsFile,
		"request_id_handler_with_group.go": withGroupFile,
	}
	for name, content := range files {
		data := ffhash.InjectCheckedLine([]byte(content))
		if err := os.WriteFile(filepath.Join(outDir, name), data, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", name, err)
		}
	}
	return nil
}
