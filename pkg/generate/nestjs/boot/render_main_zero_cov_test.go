//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderMain_ZeroCov(t *testing.T) {
	if _, err := RenderMain(nil); err == nil {
		t.Error("expected error for nil plan")
	}
	plan := &ir.BootPlan{
		ProjectID: "myapp",
		ActiveBlocks: []ir.BootBlock{
			{Name: "cors", Active: true, Config: &ir.CORSBootConfig{AllowOrigins: []string{"https://a.com"}}},
		},
	}
	out, err := RenderMain(plan)
	if err != nil {
		t.Fatalf("RenderMain error: %v", err)
	}
	for _, want := range []string{"NestFactory", "bootstrap()", "myapp", "enableCors"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderMain missing %q", want)
		}
	}
}
