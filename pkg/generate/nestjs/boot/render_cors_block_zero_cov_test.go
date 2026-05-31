//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderCorsBlock_ZeroCov(t *testing.T) {
	// No origins → simple
	var b1 strings.Builder
	renderCorsBlock(&b1, &ir.BootPlan{})
	if !strings.Contains(b1.String(), "credentials: true") {
		t.Errorf("simple cors missing:\n%s", b1.String())
	}
	// With origins
	var b2 strings.Builder
	plan := &ir.BootPlan{ActiveBlocks: []ir.BootBlock{
		{Name: "cors", Active: true, Config: &ir.CORSBootConfig{
			AllowOrigins: []string{"https://a.com"}, AllowCredentials: true,
		}},
	}}
	renderCorsBlock(&b2, plan)
	out := b2.String()
	if !strings.Contains(out, "origin: [") || !strings.Contains(out, "https://a.com") {
		t.Errorf("origin cors missing:\n%s", out)
	}
}
