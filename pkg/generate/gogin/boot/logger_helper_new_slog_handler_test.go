//ff:func feature=gen-gogin type=test control=sequence
//ff:what loggerHelperNewSlogHandler — LOG_FORMAT 기준 JSON/Text slog.Handler 생성 헬퍼 소스 반환

package boot

import (
	"strings"
	"testing"
)

func TestLoggerHelperNewSlogHandler(t *testing.T) {
	src := loggerHelperNewSlogHandler()
	for _, must := range []string{
		"func newSlogHandler(level slog.Level, keys map[string]bool) slog.Handler {",
		"redact.ReplaceAttr(keys)",
		`strings.EqualFold(os.Getenv("LOG_FORMAT"), "TEXT")`,
		"slog.NewTextHandler(os.Stdout, opts)",
		"slog.NewJSONHandler(os.Stdout, opts)",
		"&requestIDHandler{Handler: base}",
	} {
		if !strings.Contains(src, must) {
			t.Errorf("newSlogHandler helper missing %q, got:\n%s", must, src)
		}
	}
}
