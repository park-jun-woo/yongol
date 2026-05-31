//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov2 — RenderService/RenderController 진입점으로 write* 함수 0% 커버
package ssac

import (
	"strings"
	"testing"
)

func TestRenderController_Subscribe(t *testing.T) {
	out, err := RenderController(subscribeServicePlan())
	if err != nil {
		t.Fatalf("RenderController: %v", err)
	}
	if !strings.Contains(out, "Subscribe handler for topic: course.created") {
		t.Errorf("subscribe handler missing\n%s", out)
	}
	if !strings.Contains(out, "async handleOnCourseCreated(payload: any)") {
		t.Errorf("handler signature missing\n%s", out)
	}
}
