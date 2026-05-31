//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderMain_ZeroCov(t *testing.T) {
	if _, err := RenderMain(nil, nil); err == nil {
		t.Error("expected error for nil plan")
	}
	plan := &ir.BootPlan{
		ProjectID:    "myapp",
		ActiveBlocks: []ir.BootBlock{{Name: "cors", Active: true}},
	}
	out, err := RenderMain(plan, []string{"users"})
	if err != nil {
		t.Fatalf("RenderMain error: %v", err)
	}
	for _, want := range []string{"FastAPI", "myapp", "CORSMiddleware", "users_router", "/health"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderMain missing %q", want)
		}
	}
}
