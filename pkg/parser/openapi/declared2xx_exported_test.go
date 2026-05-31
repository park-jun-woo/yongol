//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestOpenAPIHelpers — unit tests for the pure openapi parser helper functions
package openapi

import (
	"testing"
)

func TestDeclared2xxExported(t *testing.T) {
	set := Declared2xx(opWith2xx(200))
	if !set[200] {
		t.Errorf("Declared2xx missing 200: %v", set)
	}
}
