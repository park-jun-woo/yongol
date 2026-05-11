//ff:func feature=stml-parse type=test control=sequence
//ff:what 페이지 data-layout 속성 추출 테스트

package stml

import (
	"strings"
	"testing"
)

func TestParsePageWithDataLayout(t *testing.T) {
	input := `<main data-layout="auth">
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <input data-field="Password" type="password" />
    <button type="submit">Login</button>
  </div>
</main>`

	page, diags := ParseReader("login.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if page.Layout != "auth" {
		t.Errorf("Layout = %q, want %q", page.Layout, "auth")
	}
}

func TestParsePageWithoutDataLayout(t *testing.T) {
	input := `<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">Login</button>
  </div>
</main>`

	page, diags := ParseReader("login.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if page.Layout != "" {
		t.Errorf("Layout = %q, want empty string", page.Layout)
	}
}
