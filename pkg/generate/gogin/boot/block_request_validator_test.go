//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what blockRequestValidator — request_validator 미들웨어 등록 (CORS 이후, Health 이전)
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRequestValidator(t *testing.T) {
	block := blockRequestValidator(&yongol.Fullstack{}, "example.com/zenflow")
	if block.Name != "request-validator" {
		t.Errorf("name = %q, want request-validator", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		"validator, err := middleware.RequestValidator()",
		"if err != nil {",
		`slog.Error("bootstrap failed", "stage", "request-validator", "err", err)`,
		"os.Exit(1)",
		"r.Use(validator)",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("blockRequestValidator missing %q, got:\n%s", must, body)
		}
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), `"example.com/zenflow/internal/middleware"`) {
		t.Errorf("must import middleware, got:\n%v", block.Imports)
	}
}
