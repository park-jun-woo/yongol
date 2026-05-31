//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestHurlOpenAPIHelpers — unit tests for the pure hurl_openapi helper functions
package hurl_openapi

import (
	"reflect"
	"testing"
)

func TestCollectNon5xxCodes(t *testing.T) {
	got := collectNon5xxCodes(map[string]bool{"200": true, "404": true, "500": true, "503": true})
	want := []string{"200", "404"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("collectNon5xxCodes = %v, want %v", got, want)
	}
	if got := collectNon5xxCodes(map[string]bool{"500": true}); got != nil {
		t.Errorf("only-5xx → %v, want nil", got)
	}
}
