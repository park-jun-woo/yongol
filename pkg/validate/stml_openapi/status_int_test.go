//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"testing"
)

func TestStatusInt(t *testing.T) {
	if got := statusInt("200"); got != 200 {
		t.Errorf("statusInt(200) = %d", got)
	}
	if got := statusInt("201"); got != 201 {
		t.Errorf("statusInt(201) = %d", got)
	}
	if got := statusInt("404"); got != 0 {
		t.Errorf("statusInt(404) = %d, want 0", got)
	}
}
