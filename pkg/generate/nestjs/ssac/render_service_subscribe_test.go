//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov2 — RenderService/RenderController 진입점으로 write* 함수 0% 커버
package ssac

import (
	"strings"
	"testing"
)

func TestRenderService_Subscribe(t *testing.T) {
	out, err := RenderService(subscribeServicePlan(), nil)
	if err != nil {
		t.Fatalf("RenderService: %v", err)
	}
	if !strings.Contains(out, "const message = payload;") {
		t.Errorf("subscribe alias missing\n%s", out)
	}
}
