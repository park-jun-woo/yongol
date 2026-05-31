//ff:func feature=crosscheck type=test control=sequence topic=scenario-check
//ff:what TestHurlHelpers — unit tests for the pure hurl parser helper functions
package hurl

import (
	"testing"
)

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
