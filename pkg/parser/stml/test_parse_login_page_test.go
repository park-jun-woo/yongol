//ff:func feature=stml-parse type=parser control=sequence
//ff:what TestParseLoginPage — login page with single action and two fields

package stml

import (
	"strings"
	"testing"
)

func TestParseLoginPage(t *testing.T) {
	input := `<main>
  <div data-action="Login" class="space-y-4">
    <input data-field="Email" type="email" />
    <input data-field="Password" type="password" />
    <button type="submit">로그인</button>
  </div>
</main>`

	page, diags := ParseReader("login-page.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	if page.Name != "login-page" {
		t.Errorf("Name = %q, want %q", page.Name, "login-page")
	}
	if len(page.Fetches) != 0 {
		t.Errorf("Fetches = %d, want 0", len(page.Fetches))
	}
	if len(page.Actions) != 1 {
		t.Fatalf("Actions = %d, want 1", len(page.Actions))
	}

	action := page.Actions[0]
	if action.OperationID != "Login" {
		t.Errorf("OperationID = %q, want %q", action.OperationID, "Login")
	}
	if len(action.Fields) != 2 {
		t.Fatalf("Fields = %d, want 2", len(action.Fields))
	}
	assertField(t, action.Fields[0], "Email", "input", "email")
	assertField(t, action.Fields[1], "Password", "input", "password")
}
