//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestWriteModuleDecorator — @Module 데코레이터 imports/controllers/providers/exports 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteModuleDecorator(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{{OperationID: "CreateCourse"}}
	deps := moduleDeps{
		NeedsQueue:           true,
		NeedsAuthz:           true,
		NeedsSameFeatureStub: true,
		CrossFeatures:        []string{"billing"},
	}
	writeModuleDecorator(&b, plans, deps, "CourseStubService")

	out := b.String()
	for _, want := range []string{
		"@Module({\n",
		"  imports: [\n",
		"    PrismaModule,\n",
		"    QueueModule,\n",
		"    AuthzModule,\n",
		"    BillingModule,\n",
		"  controllers: [\n",
		"    CreateCourseController,\n",
		"  providers: [\n",
		"    CourseStubService,\n",
		"  exports: [\n",
		"})\n",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("decorator missing %q\n--- got ---\n%s", want, out)
		}
	}
}
