//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockLoggerInit — slog 기본 핸들러 초기화 (JSON/Text, LOG_LEVEL, redact)

package boot

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockLoggerInit produces the slog default handler initialization block.
// Always active — runs first so all subsequent blocks can use slog.
// LOG_LEVEL (DEBUG/INFO/WARN/ERROR) and LOG_FORMAT (JSON/TEXT) env vars
// control output. Defaults: INFO level, JSON format.
//
// A redact.ReplaceAttr callback is installed on the handler so any slog
// attribute whose key matches redact.DefaultKeys or a DDL `-- @sensitive`
// column is rendered as "[REDACTED]". This is a defensive net for ad-hoc
// slog calls; generated sqlc row structs that contain sensitive columns
// also get a LogValue method (see pkg/generate/gogin/sqlc/) that masks
// those fields when the struct is logged whole.
func blockLoggerInit(fs *yongol.Fullstack) MainBlock {
	extras := collectSensitiveKeys(fs)
	lines := []string{
		`logLevel := parseLogLevel(os.Getenv("LOG_LEVEL"))`,
		`sensitiveKeys := buildSensitiveKeys(` + goStringSlice(extras) + `)`,
		`handler := newSlogHandler(logLevel, sensitiveKeys)`,
		`slog.SetDefault(slog.New(handler))`,
	}
	return MainBlock{
		Name: "logger-init",
		Imports: []string{
			`"log/slog"`,
			`"os"`,
			`"strings"`,
			`"github.com/park-jun-woo/ssac/pkg/redact"`,
		},
		Lines: lines,
		Funcs: []string{
			loggerHelperParseLogLevel(),
			loggerHelperBuildSensitiveKeys(),
			loggerHelperNewSlogHandler(),
		},
	}
}
