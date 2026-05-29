//ff:func feature=openapi-parse type=test control=sequence
//ff:what mapValue — nil/scalar 입력에서 nil 반환 (panic 금지)

package openapi

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMapValue_NilAndNonMapping(t *testing.T) {
	if got := mapValue(nil, "x"); got != nil {
		t.Errorf("nil in, got %v", got)
	}
	scalar := &yaml.Node{Kind: yaml.ScalarNode}
	if got := mapValue(scalar, "x"); got != nil {
		t.Errorf("scalar in, got %v", got)
	}
}
