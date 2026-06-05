//ff:func feature=stml-gen type=test control=sequence
//ff:what 액션 폼 submit 버튼 data-enabled-when을 기존 disabled에 || !(...)로 병합 검증

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionFormEnabledWhenMerge(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "CreateWorkflow",
		EnabledWhen: "workflow.status=draft",
		Fields:      []stmlparser.FieldBind{{Name: "title"}},
	}
	got := renderActionForm(a, 0)
	want := `disabled={createWorkflowMutation.isPending || !(data.workflow?.status === 'draft')}`
	if !strings.Contains(got, want) {
		t.Errorf("form merge:\n  got:  %s\n  want substring: %s", got, want)
	}
}
