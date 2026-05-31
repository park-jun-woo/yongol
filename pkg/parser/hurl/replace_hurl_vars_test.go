//ff:func feature=crosscheck type=test control=sequence topic=scenario-check
//ff:what TestHurlHelpers — unit tests for the pure hurl parser helper functions
package hurl

import (
	"testing"
)

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
