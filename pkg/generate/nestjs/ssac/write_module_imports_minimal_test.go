//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteModuleImportsMinimal — TestWriteModuleImports — NestJS module 파일 상단 import 블록 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteModuleImportsMinimal(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{{OperationID: "ListCourses"}}
	writeModuleImports(&b, "Course", plans, moduleDeps{}, "Unused")

	out := b.String()
	if strings.Contains(out, "QueueModule") || strings.Contains(out, "AuthzModule") {
		t.Errorf("minimal deps should not import Queue/Authz, got %q", out)
	}
	if strings.Contains(out, "Unused") {
		t.Errorf("stub service should not appear without NeedsSameFeatureStub, got %q", out)
	}
}
