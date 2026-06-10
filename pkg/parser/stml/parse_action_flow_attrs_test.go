//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseActionFlowAttrs — data-capture/data-redirect/data-on-error가 ActionBlock에 파싱되는지 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseActionFlowAttrs(t *testing.T) {
	input := `<main>
  <section data-action="Login"
           data-capture="access_token -> auth.token, refresh_token -> auth.refresh"
           data-redirect="/">
    <input data-field="email" type="email" />
    <input data-field="password" type="password" />
    <button type="submit">로그인</button>
    <p data-on-error></p>
  </section>
</main>`

	page, diags := ParseReader("login.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if len(page.Actions) != 1 {
		t.Fatalf("Actions = %d, want 1", len(page.Actions))
	}
	a := page.Actions[0]
	if a.CaptureRaw != "access_token -> auth.token, refresh_token -> auth.refresh" {
		t.Errorf("CaptureRaw = %q", a.CaptureRaw)
	}
	if len(a.Captures) != 2 {
		t.Fatalf("Captures = %d, want 2: %+v", len(a.Captures), a.Captures)
	}
	if a.Captures[0] != (CaptureBind{RespField: "access_token", Sink: "auth.token"}) {
		t.Errorf("Captures[0] = %+v", a.Captures[0])
	}
	if a.Captures[1] != (CaptureBind{RespField: "refresh_token", Sink: "auth.refresh"}) {
		t.Errorf("Captures[1] = %+v", a.Captures[1])
	}
	if a.Redirect != "/" {
		t.Errorf("Redirect = %q, want %q", a.Redirect, "/")
	}
	if !a.OnErrorNode {
		t.Errorf("OnErrorNode = false, want true")
	}
	if len(page.FlowAttrMisplaced) != 0 {
		t.Errorf("FlowAttrMisplaced = %+v, want none", page.FlowAttrMisplaced)
	}

	// Invalid capture syntax: raw is kept, parsed bindings stay empty
	// (TM-20 reports the syntax error at validate time).
	input = `<main>
  <section data-action="Login" data-capture="access_token => session.token">
    <button type="submit">go</button>
  </section>
</main>`
	page, diags = ParseReader("login.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	a = page.Actions[0]
	if a.CaptureRaw != "access_token => session.token" {
		t.Errorf("invalid: CaptureRaw = %q", a.CaptureRaw)
	}
	if len(a.Captures) != 0 {
		t.Errorf("invalid: Captures = %+v, want none", a.Captures)
	}
	if a.OnErrorNode {
		t.Errorf("invalid: OnErrorNode = true, want false")
	}
}
