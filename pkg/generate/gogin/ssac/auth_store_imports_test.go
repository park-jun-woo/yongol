//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what authStoreImports 단위 테스트 (subscribe/wrap_calls/gin 의존 import 조립)

package ssac

import (
	"strings"
	"testing"
)

func contains(s []string, want string) bool {
	for _, v := range s {
		if v == want {
			return true
		}
	}
	return false
}

func TestAuthStoreImports(t *testing.T) {
	t.Run("http RefreshRotate includes slog + gin", func(t *testing.T) {
		g := &methodGen{}
		imps := g.authStoreImports("auth", "RefreshRotate")
		if !contains(imps, `"github.com/park-jun-woo/ssac/pkg/auth"`) {
			t.Errorf("missing auth pkg import: %v", imps)
		}
		if !contains(imps, `"log/slog"`) {
			t.Errorf("http path should import log/slog: %v", imps)
		}
		if !contains(imps, `"github.com/gin-gonic/gin"`) {
			t.Errorf("RefreshRotate should import gin: %v", imps)
		}
		if contains(imps, `"fmt"`) {
			t.Errorf("http path should not import fmt: %v", imps)
		}
	})
	t.Run("subscribe imports fmt not slog/gin", func(t *testing.T) {
		g := &methodGen{IsSubscribe: true}
		imps := g.authStoreImports("auth", "RefreshRotate")
		if !contains(imps, `"fmt"`) {
			t.Errorf("subscribe should import fmt: %v", imps)
		}
		if contains(imps, `"log/slog"`) || contains(imps, `"github.com/gin-gonic/gin"`) {
			t.Errorf("subscribe should not import slog/gin: %v", imps)
		}
	})
	t.Run("wrap_calls adds otel", func(t *testing.T) {
		g := &methodGen{WrapCalls: true}
		imps := g.authStoreImports("auth", "Logout")
		if !contains(imps, `"go.opentelemetry.io/otel"`) {
			t.Errorf("WrapCalls should add otel: %v", imps)
		}
	})
	t.Run("pkg name interpolated", func(t *testing.T) {
		g := &methodGen{}
		imps := g.authStoreImports("session", "Other")
		joined := strings.Join(imps, " ")
		if !strings.Contains(joined, "ssac/pkg/session") {
			t.Errorf("expected session pkg import, got %v", imps)
		}
		// non Refresh/Logout call → no gin
		if contains(imps, `"github.com/gin-gonic/gin"`) {
			t.Errorf("non-Refresh/Logout should not import gin: %v", imps)
		}
	})
}
