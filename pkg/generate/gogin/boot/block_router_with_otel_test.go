//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockRouter — gin.Default + (OTel 활성 시) otelgin 미들웨어 등록
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockRouter_WithOtel(t *testing.T) {
	fs := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true})
	block := blockRouter(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "r.Use(otelgin.Middleware(otelServiceName))") {
		t.Errorf("otel project must register otelgin, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), "otelgin") {
		t.Errorf("must import otelgin, got:\n%v", block.Imports)
	}
}
