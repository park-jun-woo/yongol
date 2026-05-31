//ff:func feature=gen-gogin type=test control=sequence
//ff:what authStoreErrBranch 단위 테스트 (subscribe fmt.Errorf vs http log+JSONResponse)
package ssac

import (
	"strings"
	"testing"
)

func TestAuthStoreErrBranch(t *testing.T) {
	t.Run("http branch logs + returns JSON envelope", func(t *testing.T) {
		g := &methodGen{FuncName: "Login"}
		lines := g.authStoreErrBranch("auth", "RefreshRotate", 401, "Unauthorized")
		joined := strings.Join(lines, "\n")
		if lines[0] != "if err != nil {" {
			t.Errorf("first line = %q", lines[0])
		}
		if !strings.Contains(joined, "slog.Warn") {
			t.Errorf("401 should use slog.Warn:\n%s", joined)
		}
		if !strings.Contains(joined, `api.Login401JSONResponse{Error: "Unauthorized", Code: "unauthorized"}, nil`) {
			t.Errorf("missing JSON envelope:\n%s", joined)
		}
		if lines[len(lines)-1] != "}" {
			t.Errorf("last line = %q, want }", lines[len(lines)-1])
		}
	})
	t.Run("subscribe wraps error", func(t *testing.T) {
		g := &methodGen{IsSubscribe: true}
		lines := g.authStoreErrBranch("auth", "Logout", 500, "boom")
		joined := strings.Join(lines, "\n")
		if !strings.Contains(joined, `return fmt.Errorf("auth.Logout: %w", err)`) {
			t.Errorf("missing wrapped error:\n%s", joined)
		}
	})
}
