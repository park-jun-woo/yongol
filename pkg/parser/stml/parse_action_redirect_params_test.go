//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseActionRedirectParams — data-redirect 페이지명 참조 + data-redirect-params가 ActionBlock에 파싱되는지 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseActionRedirectParams(t *testing.T) {
	input := `<main>
  <div data-action="CreateContract"
       data-redirect="contract-edit"
       data-redirect-params="id -> ContractID">
    <input data-field="Title" type="text" />
    <button type="submit">등록</button>
  </div>
</main>`

	page, diags := ParseReader("contract-new.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if len(page.Actions) != 1 {
		t.Fatalf("Actions = %d, want 1", len(page.Actions))
	}
	a := page.Actions[0]
	if a.Redirect != "contract-edit" {
		t.Errorf("Redirect = %q, want %q", a.Redirect, "contract-edit")
	}
	if a.RedirectParamsRaw != "id -> ContractID" {
		t.Errorf("RedirectParamsRaw = %q", a.RedirectParamsRaw)
	}
	if len(a.RedirectParams) != 1 || a.RedirectParams[0] != (LinkParamBind{Source: "id", Segment: "ContractID"}) {
		t.Errorf("RedirectParams = %+v", a.RedirectParams)
	}
	if a.RedirectPattern != "" {
		t.Errorf("RedirectPattern = %q, want empty (codegen-populated)", a.RedirectPattern)
	}

	// Invalid syntax: raw is kept, parsed bindings stay empty (TM-33
	// reports the syntax error at validate time).
	input = `<main>
  <div data-action="CreateContract" data-redirect="contract-edit" data-redirect-params="item.id -> ContractID">
    <button type="submit">go</button>
  </div>
</main>`
	page, diags = ParseReader("contract-new.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	a = page.Actions[0]
	if a.RedirectParamsRaw != "item.id -> ContractID" {
		t.Errorf("invalid: RedirectParamsRaw = %q", a.RedirectParamsRaw)
	}
	if len(a.RedirectParams) != 0 {
		t.Errorf("invalid: RedirectParams = %+v, want none", a.RedirectParams)
	}
}
