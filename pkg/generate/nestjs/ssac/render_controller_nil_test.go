//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov2 — RenderService/RenderController 진입점으로 write* 함수 0% 커버
package ssac

import (
	"testing"
)

func TestRenderController_Nil(t *testing.T) {
	if _, err := RenderController(nil); err == nil {
		t.Error("nil plan should error")
	}
}
