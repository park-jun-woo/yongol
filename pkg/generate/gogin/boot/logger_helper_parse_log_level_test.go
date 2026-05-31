//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what loggerHelperParseLogLevel — LOG_LEVEL 문자열 → slog.Level 변환 헬퍼 소스 반환
package boot

import (
	"strings"
	"testing"
)

func TestLoggerHelperParseLogLevel(t *testing.T) {
	src := loggerHelperParseLogLevel()
	for _, must := range []string{
		"func parseLogLevel(level string) slog.Level {",
		"strings.ToUpper(level)",
		`case "DEBUG":`,
		"slog.LevelDebug",
		`case "WARN":`,
		"slog.LevelWarn",
		`case "ERROR":`,
		"slog.LevelError",
		"return slog.LevelInfo",
	} {
		if !strings.Contains(src, must) {
			t.Errorf("parseLogLevel helper missing %q, got:\n%s", must, src)
		}
	}
}
