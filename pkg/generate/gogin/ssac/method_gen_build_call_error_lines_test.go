//ff:func feature=gen-gogin type=test control=sequence
//ff:what buildCallErrorLines 단위 테스트 (HTTP log+JSON vs subscribe fmt.Errorf)
package ssac

import (
	"strings"
	"testing"
)

func TestBuildCallErrorLines(t *testing.T) {
	t.Run("http branch", func(t *testing.T) {
		g := &methodGen{FuncName: "RunReport"}
		lines := g.buildCallErrorLines("dashboard", "Summarize", "Report failed", 500)
		joined := strings.Join(lines, "\n")
		if lines[0] != "if err != nil {" || lines[len(lines)-1] != "}" {
			t.Errorf("expected if/} wrapper, got %v", lines)
		}
		if !strings.Contains(joined, "slog.Error") {
			t.Errorf("500 should log via slog.Error:\n%s", joined)
		}
		if !strings.Contains(joined, `api.RunReport500JSONResponse{Error: "Report failed", Code: "internal_error"}, nil`) {
			t.Errorf("missing JSON envelope:\n%s", joined)
		}
	})
	t.Run("subscribe branch wraps error", func(t *testing.T) {
		g := &methodGen{IsSubscribe: true}
		lines := g.buildCallErrorLines("queue", "Drain", "msg", 500)
		joined := strings.Join(lines, "\n")
		if !strings.Contains(joined, `return fmt.Errorf("queue.Drain: %w", err)`) {
			t.Errorf("missing wrapped error:\n%s", joined)
		}
	})
}
