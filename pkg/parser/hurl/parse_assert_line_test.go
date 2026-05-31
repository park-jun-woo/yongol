//ff:func feature=crosscheck type=test control=sequence topic=scenario-check
//ff:what TestHurlHelpers — unit tests for the pure hurl parser helper functions
package hurl

import (
	"testing"
)

func TestParseAssertLine(t *testing.T) {
	a, ok := parseAssertLine(`jsonpath "$.user.id" isInteger`, 7)
	if !ok || a.JSONPath != "$.user.id" || a.Line != 7 {
		t.Errorf("assert = %+v ok=%v", a, ok)
	}
	if _, ok := parseAssertLine(`status == 200`, 1); ok {
		t.Error("non-jsonpath line should not parse as assert")
	}
}
