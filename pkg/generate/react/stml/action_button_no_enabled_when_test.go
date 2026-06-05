//ff:func feature=stml-gen type=test control=sequence
//ff:what data-enabled-when 없는 액션 버튼 disabled 출력이 불변임을 검증 (회귀)

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionButtonNoEnabledWhenUnchanged(t *testing.T) {
	a := stmlparser.ActionBlock{OperationID: "ActivateWorkflow"}
	got := renderActionButton(a, 0, nil)
	want := `disabled={activateWorkflowMutation.isPending}`
	if !strings.Contains(got, want) {
		t.Errorf("button no-enabled-when:\n  got: %s\n  want substring: %s", got, want)
	}
}
