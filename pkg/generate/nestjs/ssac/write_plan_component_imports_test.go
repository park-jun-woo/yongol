//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWritePlanComponentImports — TestWritePlanComponentImports — plan 별 Controller/Service import 문 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWritePlanComponentImports(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{{OperationID: "CreateCourse"}}
	writePlanComponentImports(&b, plans)

	want := "import { CreateCourseController } from './createCourse.controller';\n" +
		"import { CreateCourseService } from './createCourse.service';\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
