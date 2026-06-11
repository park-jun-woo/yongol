//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseActionPrefill — data-prefill 값이 ActionBlock.Prefill에 파싱되는지 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseActionPrefill(t *testing.T) {
	input := `<main>
  <article data-fetch="GetRule" data-param-rule-id="route.RuleID">
    <form data-action="UpdateRule" data-prefill="GetRule" data-param-rule-id="route.RuleID">
      <input data-field="sheet_name" />
      <button type="submit">save</button>
    </form>
  </article>
</main>`

	page, diags := ParseReader("rule-edit.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	actions := CollectChildActions(page.Children)
	if len(actions) != 1 {
		t.Fatalf("actions = %d, want 1", len(actions))
	}
	if actions[0].Prefill != "GetRule" {
		t.Errorf("Prefill = %q, want %q", actions[0].Prefill, "GetRule")
	}
	if len(page.FlowAttrMisplaced) != 0 {
		t.Errorf("FlowAttrMisplaced = %+v, want none (data-prefill on the action element is valid)", page.FlowAttrMisplaced)
	}

	// Absent data-prefill leaves Prefill empty.
	input = `<main><form data-action="CreateRule"><button type="submit">go</button></form></main>`
	page, _ = ParseReader("rule-new.html", strings.NewReader(input))
	if page.Actions[0].Prefill != "" {
		t.Errorf("Prefill = %q, want empty", page.Actions[0].Prefill)
	}
}
