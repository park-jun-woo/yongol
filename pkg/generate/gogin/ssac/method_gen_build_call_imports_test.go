//ff:func feature=gen-gogin type=test control=sequence
//ff:what buildCallImports 단위 테스트 (ImportMap/auth model/slog/otel/gin 분기)
package ssac

import (
	"testing"
)

func TestBuildCallImports(t *testing.T) {
	t.Run("project pkg import + slog for http", func(t *testing.T) {
		g := &methodGen{
			ImportMap: map[string]string{"dashboard": "github.com/x/internal/dashboard"},
		}
		imps := g.buildCallImports("dashboard", "Summarize", "res")
		if !contains(imps, `"github.com/x/internal/dashboard"`) {
			t.Errorf("missing dashboard import: %v", imps)
		}
		if !contains(imps, `"log/slog"`) {
			t.Errorf("http should import log/slog: %v", imps)
		}
	})
	t.Run("auth IssueToken adds model import", func(t *testing.T) {
		g := &methodGen{ModulePath: "github.com/x"}
		imps := g.buildCallImports("auth", "IssueToken", "tok")
		if !contains(imps, `"github.com/x/internal/model"`) {
			t.Errorf("auth IssueToken should import model: %v", imps)
		}
	})
	t.Run("subscribe imports fmt + otel when wrapped", func(t *testing.T) {
		g := &methodGen{IsSubscribe: true, WrapCalls: true}
		imps := g.buildCallImports("queue", "Drain", "_")
		if !contains(imps, `"fmt"`) {
			t.Errorf("subscribe should import fmt: %v", imps)
		}
		if !contains(imps, `"go.opentelemetry.io/otel"`) {
			t.Errorf("WrapCalls should import otel: %v", imps)
		}
	})
	t.Run("RefreshToken with access var adds gin", func(t *testing.T) {
		g := &methodGen{AccessTokenVar: "issued", ModulePath: "github.com/x"}
		imps := g.buildCallImports("auth", "RefreshToken", "refreshed")
		if !contains(imps, `"github.com/gin-gonic/gin"`) {
			t.Errorf("RefreshToken should import gin: %v", imps)
		}
	})
}
