//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestHurlOpenAPIHelpers — unit tests for the pure hurl_openapi helper functions
package hurl_openapi

import (
	"reflect"
	"testing"
)

func TestFilterUncovered(t *testing.T) {
	declared := []string{"200", "404", "409"}
	covered := map[string]bool{"200": true}
	got := filterUncovered(declared, covered)
	want := []string{"404", "409"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterUncovered = %v, want %v", got, want)
	}
	// nil coveredSet → all declared are uncovered.
	if got := filterUncovered(declared, nil); !reflect.DeepEqual(got, declared) {
		t.Errorf("nil covered → %v, want %v", got, declared)
	}
}
