//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseActionEnabledWhen — data-enabled-when 속성 추출 + 미설정 시 빈 값

package stml

import (
	"strings"
	"testing"
)

func TestParseActionEnabledWhen(t *testing.T) {
	input := `<main>
  <button data-action="ActivateWorkflow" data-param-id="route.WorkflowID" data-enabled-when="workflow.status = draft">활성화</button>
  <button data-action="DeleteWorkflow" data-param-id="route.WorkflowID">삭제</button>
</main>`

	page, diags := ParseReader("test.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if len(page.Actions) != 2 {
		t.Fatalf("page.Actions = %d, want 2", len(page.Actions))
	}
	if got := page.Actions[0].EnabledWhen; got != "workflow.status = draft" {
		t.Errorf("Actions[0].EnabledWhen = %q, want %q", got, "workflow.status = draft")
	}
	if got := page.Actions[1].EnabledWhen; got != "" {
		t.Errorf("Actions[1].EnabledWhen = %q, want empty", got)
	}
}
