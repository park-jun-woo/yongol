//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteModuleExportsSectionNoStub — TestWriteModuleExportsSection — @Module exports 배열(서비스 + stub) 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteModuleExportsSectionNoStub(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{{OperationID: "CreateCourse"}}
	writeModuleExportsSection(&b, plans, moduleDeps{}, "Unused")

	want := "  exports: [\n    CreateCourseService,\n  ],\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
