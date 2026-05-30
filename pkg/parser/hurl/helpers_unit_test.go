//ff:func feature=crosscheck type=test control=sequence topic=scenario-check
//ff:what TestHurlHelpers — unit tests for the pure hurl parser helper functions
package hurl

import (
	"sort"
	"testing"
)

func TestTrimQuery(t *testing.T) {
	tests := []struct{ in, want string }{
		{"/users?page=2", "/users"},
		{"/users", "/users"},
		{"/a?b?c", "/a"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := trimQuery(tt.in); got != tt.want {
			t.Errorf("trimQuery(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestReplaceHurlVars(t *testing.T) {
	// Bare placeholder value gets quoted.
	if got := replaceHurlVars(`{"id": {{userID}}}`); got != `{"id": "__hurl_var__"}` {
		t.Errorf("bare value = %q", got)
	}
	// Placeholder already inside quotes stays bare.
	if got := replaceHurlVars(`{"email": "{{email}}"}`); got != `{"email": "__hurl_var__"}` {
		t.Errorf("quoted value = %q", got)
	}
	// No placeholder → unchanged.
	if got := replaceHurlVars(`{"a": 1}`); got != `{"a": 1}` {
		t.Errorf("no placeholder = %q", got)
	}
	// Unterminated placeholder → appended verbatim.
	if got := replaceHurlVars(`{{unterminated`); got != `{{unterminated` {
		t.Errorf("unterminated = %q", got)
	}
}

func TestExtractJSONFieldNames(t *testing.T) {
	got := extractJSONFieldNames(`{"id": 1, "name": "x"}`)
	sort.Strings(got)
	if len(got) != 2 || got[0] != "id" || got[1] != "name" {
		t.Errorf("fields = %v, want [id name]", got)
	}
	// Body with placeholders still yields keys.
	got2 := extractJSONFieldNames(`{"email": "{{email}}", "n": {{n}}}`)
	sort.Strings(got2)
	if len(got2) != 2 {
		t.Errorf("placeholder fields = %v", got2)
	}
	// Non-object → nil.
	if got := extractJSONFieldNames(`[1,2,3]`); got != nil {
		t.Errorf("array → %v, want nil", got)
	}
	// Malformed → nil.
	if got := extractJSONFieldNames(`{not json`); got != nil {
		t.Errorf("malformed → %v, want nil", got)
	}
}

func TestParseAssertLine(t *testing.T) {
	a, ok := parseAssertLine(`jsonpath "$.user.id" isInteger`, 7)
	if !ok || a.JSONPath != "$.user.id" || a.Line != 7 {
		t.Errorf("assert = %+v ok=%v", a, ok)
	}
	if _, ok := parseAssertLine(`status == 200`, 1); ok {
		t.Error("non-jsonpath line should not parse as assert")
	}
}

func TestParseCaptureLine(t *testing.T) {
	// jsonpath form.
	c, ok := parseCaptureLine(`token: jsonpath "$.access_token"`, 3)
	if !ok || c.Var != "token" || c.Source != "jsonpath" || c.JSONPath != "$.access_token" {
		t.Errorf("jsonpath capture = %+v ok=%v", c, ok)
	}
	// header form.
	c2, ok := parseCaptureLine(`loc: header "Location"`, 4)
	if !ok || c2.Source != "header" || c2.Header != "Location" {
		t.Errorf("header capture = %+v ok=%v", c2, ok)
	}
	// raw expression stored verbatim.
	c3, ok := parseCaptureLine(`x: cookie "session"`, 5)
	if !ok || c3.Source != `cookie "session"` {
		t.Errorf("raw capture = %+v ok=%v", c3, ok)
	}
	// non-capture line.
	if _, ok := parseCaptureLine(`GET https://example.com`, 1); ok {
		t.Error("request line should not parse as capture")
	}
}
