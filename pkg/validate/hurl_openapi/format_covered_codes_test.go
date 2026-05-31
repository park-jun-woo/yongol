//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestHurlOpenAPIHelpers — unit tests for the pure hurl_openapi helper functions
package hurl_openapi

import (
	"testing"
)

func TestFormatCoveredCodes(t *testing.T) {
	if got := formatCoveredCodes(map[string]bool{"201": true, "200": true}); got != "200, 201" {
		t.Errorf("got %q, want '200, 201'", got)
	}
	if got := formatCoveredCodes(map[string]bool{}); got != "none" {
		t.Errorf("empty → %q, want none", got)
	}
}
