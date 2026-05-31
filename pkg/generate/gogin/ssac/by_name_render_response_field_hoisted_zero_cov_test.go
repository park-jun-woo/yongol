//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"strings"
	"testing"
)

func TestByNameRenderResponseFieldHoisted_ZeroCov(t *testing.T) {
	g := &methodGen{RespFields: map[string]responseField{
		"workflow": {JSONName: "workflow", RefType: "Workflow"},
		"count":    {JSONName: "count", IsRequired: true},
	}}
	scalarLocal := map[string]string{"workflow": "wfLocal"}
	listLocal := map[string]string{}

	// $ref → delegate.
	if got := g.renderResponseFieldHoisted("workflow", "x", scalarLocal, listLocal); got == "" {
		t.Errorf("ref hoisted empty")
	}
	// required non-ref → direct.
	if got := g.renderResponseFieldHoisted("count", "n", scalarLocal, listLocal); !strings.Contains(got, "Count: n,") {
		t.Errorf("required hoisted = %q", got)
	}
	// integer literal → int64 wrap + ptrOf.
	if got := g.renderResponseFieldHoisted("total", "5", scalarLocal, listLocal); !strings.Contains(got, "int64(5)") {
		t.Errorf("int literal hoisted = %q", got)
	}
	// variable → &expr.
	if got := g.renderResponseFieldHoisted("name", "v", scalarLocal, listLocal); !strings.Contains(got, "&v,") {
		t.Errorf("var hoisted = %q", got)
	}
}
