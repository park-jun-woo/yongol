//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestZeroCov2 — RenderService/RenderController 진입점으로 write* 함수 0% 커버
package ssac

import (
	"strings"
	"testing"
)

func TestRenderService_HTTP(t *testing.T) {
	out, err := RenderService(httpServicePlan(), nil)
	if err != nil {
		t.Fatalf("RenderService: %v", err)
	}
	for _, want := range []string{
		"@Injectable()",
		"export class CreateCourseService {",
		"private readonly prisma: PrismaService,",
		"private readonly queue: QueueService,",
		"private readonly authz: AuthzService,",
		"billingService",
		"$transaction(async (tx)",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderService HTTP missing %q\n%s", want, out)
		}
	}
}
