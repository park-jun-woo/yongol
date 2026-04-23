//ff:func feature=rule type=test control=sequence
//ff:what resolveSchemaType — nil 입력 시 빈 문자열 반환

package ground

import (
	"testing"
)

// TestResolveSchemaType_Nil returns "".
func TestResolveSchemaType_Nil(t *testing.T) {
	if got := resolveSchemaType(nil); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}
