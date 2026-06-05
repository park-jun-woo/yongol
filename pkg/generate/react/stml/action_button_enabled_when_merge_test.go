//ff:func feature=stml-gen type=test control=sequence
//ff:what 액션 버튼 data-enabled-when을 기존 disabled에 || !(...)로 병합 검증

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionButtonEnabledWhenMerge(t *testing.T) {
	a := stmlparser.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=draft"}
	got := renderActionButton(a, 0, nil)
	want := `disabled={activateWorkflowMutation.isPending || !(data.workflow?.status === 'draft')}`
	if !strings.Contains(got, want) {
		t.Errorf("button merge:\n  got:  %s\n  want substring: %s", got, want)
	}
}
