//ff:func feature=rule type=test control=sequence
//ff:what extractRequestFields — nil content → nil 반환

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestFields_NilContent(t *testing.T) {
	body := &openapi3.RequestBody{}
	if got := extractRequestFields(body); got != nil {
		t.Errorf("expected nil for empty body, got %v", got)
	}
}
