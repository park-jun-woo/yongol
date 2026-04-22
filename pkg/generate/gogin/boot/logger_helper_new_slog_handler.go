//ff:func feature=gen-gogin type=generator control=sequence
//ff:what loggerHelperNewSlogHandler — LOG_FORMAT 기준 JSON/Text slog.Handler 생성 헬퍼 소스 반환

package boot

// loggerHelperNewSlogHandler returns the top-level newSlogHandler(level, keys) helper
// source. Centralizes handler construction so main() stays flat (no depth-1 if at all
// in this section). LOG_FORMAT=TEXT selects slog.NewTextHandler, otherwise JSON.
//
// Phase004: wraps the base handler in requestIDHandler (defined in
// cmd/request_id_handler.go via writeRequestIDHandlerFile) so every *Context
// slog call inherits the request_id key from the context.Context created by
// the RequestID middleware.
func loggerHelperNewSlogHandler() string {
	return `func newSlogHandler(level slog.Level, keys map[string]bool) slog.Handler {
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
}`
}
