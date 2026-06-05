//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteModuleProvidersSectionWithStub — TestWriteModuleProvidersSection — @Module providers 배열(서비스 + stub) 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteModuleProvidersSectionWithStub(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{{OperationID: "CreateCourse"}}
	deps := moduleDeps{NeedsSameFeatureStub: true}
	writeModuleProvidersSection(&b, plans, deps, "CourseStubService")

	want := "  providers: [\n    CreateCourseService,\n    CourseStubService,\n  ],\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
