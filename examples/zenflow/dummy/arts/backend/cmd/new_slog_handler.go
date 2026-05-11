//ff:func feature=main type=util control=sequence
//ff:what newSlogHandler — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=36d44e46
package main

import (
	"github.com/park-jun-woo/ssac/pkg/redact"
	"log/slog"
	"os"
	"strings"
)

func newSlogHandler(level slog.Level, keys map[string]bool) slog.Handler {
	opts := &slog.HandlerOptions{
		Level:       level,
		ReplaceAttr: redact.ReplaceAttr(keys),
	}
	var base slog.Handler
	if strings.EqualFold(os.Getenv("LOG_FORMAT"), "TEXT") {
		base = slog.NewTextHandler(os.Stdout, opts)
	} else {
		base = slog.NewJSONHandler(os.Stdout, opts)
	}
	return &requestIDHandler{Handler: base}
}
