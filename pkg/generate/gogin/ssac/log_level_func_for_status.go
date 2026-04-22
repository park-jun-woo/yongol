//ff:func feature=gen-gogin type=util control=sequence
//ff:what logLevelFuncForStatus — HTTP status 4xx→slog.Warn, 그 외→slog.Error

package ssac

// logLevelFuncForStatus returns the slog function name to call:
//   - 4xx → slog.Warn
//   - else (5xx, etc.) → slog.Error
func logLevelFuncForStatus(status int) string {
	if status >= 400 && status < 500 {
		return "slog.Warn"
	}
	return "slog.Error"
}
