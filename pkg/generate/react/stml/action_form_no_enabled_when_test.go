//ff:func feature=stml-gen type=test control=sequence
//ff:what data-enabled-when 없는 액션 폼 submit 버튼 disabled 출력이 불변임을 검증 (회귀)

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionFormNoEnabledWhenUnchanged(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "CreateWorkflow",
		Fields:      []stmlparser.FieldBind{{Name: "title"}},
	}
	got := renderActionForm(a, 0)
	want := `disabled={createWorkflowMutation.isPending}`
	if !strings.Contains(got, want) {
		t.Errorf("form no-enabled-when:\n  got: %s\n  want substring: %s", got, want)
	}
}
