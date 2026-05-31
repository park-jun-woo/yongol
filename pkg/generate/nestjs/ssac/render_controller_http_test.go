//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestZeroCov2 — RenderService/RenderController 진입점으로 write* 함수 0% 커버
package ssac

import (
	"strings"
	"testing"
)

func TestRenderController_HTTP(t *testing.T) {
	out, err := RenderController(httpServicePlan())
	if err != nil {
		t.Fatalf("RenderController: %v", err)
	}
	for _, want := range []string{
		"@Controller('courses')",
		"export class CreateCourseController {",
		"from '@nestjs/common';",
		"Param,",
		"Body,",
		"Query,",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderController HTTP missing %q\n%s", want, out)
		}
	}
}
