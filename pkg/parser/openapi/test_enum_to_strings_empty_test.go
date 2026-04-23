//ff:func feature=openapi-parse type=test control=sequence
//ff:what enumToStrings — nil 입력은 nil 반환

package openapi

import "testing"

func TestEnumToStrings_Empty(t *testing.T) {
	if got := enumToStrings(nil); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}
