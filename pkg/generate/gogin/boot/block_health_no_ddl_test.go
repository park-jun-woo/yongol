//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockHealth — /health (liveness) + /ready (readiness) 등록
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockHealth_NoDDL(t *testing.T) {
	block := blockHealth(&yongol.Fullstack{})
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `r.GET("/health"`) {
		t.Errorf("must register /health, got:\n%s", body)
	}
	// no DDL → static /ready, no pgxpool helper.
	if !strings.Contains(body, `r.GET("/ready", func(c *gin.Context) {`) {
		t.Errorf("no-DDL /ready should be static, got:\n%s", body)
	}
	if len(block.Funcs) != 0 {
		t.Errorf("no-DDL health should emit no helper funcs, got %d", len(block.Funcs))
	}
	if strings.Contains(strings.Join(block.Imports, "\n"), "pgxpool") {
		t.Errorf("no-DDL health must not import pgxpool, got:\n%v", block.Imports)
	}
}
