//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteModuleExportsSectionWithStub — TestWriteModuleExportsSection — @Module exports 배열(서비스 + stub) 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteModuleExportsSectionWithStub(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{{OperationID: "CreateCourse"}}
	deps := moduleDeps{NeedsSameFeatureStub: true}
	writeModuleExportsSection(&b, plans, deps, "CourseStubService")

	want := "  exports: [\n    CreateCourseService,\n    CourseStubService,\n  ],\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
