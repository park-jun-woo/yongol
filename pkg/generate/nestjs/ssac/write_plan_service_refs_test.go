//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWritePlanServiceRefs — plan 별 Service 참조 항목 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWritePlanServiceRefs(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{
		{OperationID: "CreateCourse"},
		{OperationID: "DeleteCourse"},
	}
	writePlanServiceRefs(&b, plans)

	want := "    CreateCourseService,\n    DeleteCourseService,\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
