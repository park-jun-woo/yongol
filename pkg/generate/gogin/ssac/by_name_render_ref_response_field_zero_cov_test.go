//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"strings"
	"testing"
)

func TestByNameRenderRefResponseField_ZeroCov(t *testing.T) {
	scalarLocal := map[string]string{"workflow": "wfLocal"}
	listLocal := map[string]string{"actions": "actsLocal"}

	// scalar required.
	rf := responseField{JSONName: "workflow", RefType: "Workflow", IsRequired: true}
	got := renderRefResponseField("Workflow", "workflow", rf, scalarLocal, listLocal)
	if !strings.Contains(got, "Workflow: *wfLocal,") {
		t.Errorf("scalar required ref = %q", got)
	}
	// scalar optional.
	rfOpt := responseField{JSONName: "workflow", RefType: "Workflow"}
	gotOpt := renderRefScalarResponseField("Workflow", "workflow", rfOpt, scalarLocal)
	if !strings.Contains(gotOpt, "Workflow: wfLocal,") {
		t.Errorf("scalar optional ref = %q", gotOpt)
	}
	// array branch.
	rfArr := responseField{JSONName: "actions", RefType: "Action", IsArray: true}
	gotArr := renderRefResponseField("Actions", "actions", rfArr, scalarLocal, listLocal)
	if gotArr == "" {
		t.Errorf("array ref empty")
	}
}
