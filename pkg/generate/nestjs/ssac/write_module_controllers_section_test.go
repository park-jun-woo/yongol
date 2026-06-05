//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteModuleControllersSection — TestWriteModuleControllersSection — @Module controllers 배열 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteModuleControllersSection(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{
		{OperationID: "CreateCourse"},
		{OperationID: "DeleteCourse"},
	}
	writeModuleControllersSection(&b, plans)

	want := "  controllers: [\n    CreateCourseController,\n    DeleteCourseController,\n  ],\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
