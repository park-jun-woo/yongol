//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestParseActionInvalidates — data-invalidates 공백 분리 추출 + 미설정 시 빈 값

package stml

import (
	"strings"
	"testing"
)

func TestParseActionInvalidates(t *testing.T) {
	input := `<main>
  <div data-action="CreateWorkflow" data-invalidates="ListWorkflows  GetWorkflow">
    <input data-field="title" /><button type="submit">생성</button>
  </div>
  <button data-action="DeleteWorkflow" data-param-id="route.WorkflowID">삭제</button>
</main>`

	page, diags := ParseReader("test.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if len(page.Actions) != 2 {
		t.Fatalf("page.Actions = %d, want 2", len(page.Actions))
	}
	got := page.Actions[0].Invalidates
	want := []string{"ListWorkflows", "GetWorkflow"}
	if len(got) != len(want) {
		t.Fatalf("Actions[0].Invalidates = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Actions[0].Invalidates[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if len(page.Actions[1].Invalidates) != 0 {
		t.Errorf("Actions[1].Invalidates = %v, want empty", page.Actions[1].Invalidates)
	}
}
