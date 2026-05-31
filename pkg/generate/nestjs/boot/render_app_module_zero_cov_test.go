//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"strings"
	"testing"
)

func TestRenderAppModule_ZeroCov(t *testing.T) {
	out, err := RenderAppModule([]string{"billing"}, []string{"queue"})
	if err != nil {
		t.Fatalf("RenderAppModule error: %v", err)
	}
	for _, want := range []string{"PrismaModule", "QueueModule", "BillingModule", "export class AppModule"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderAppModule missing %q", want)
		}
	}
}
