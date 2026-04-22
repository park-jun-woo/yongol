//ff:func feature=validate type=test control=sequence topic=rego-structural
//ff:what XPP-30 테스트 (TODO: 케이스 추가)

package rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXpp30OwnershipNoAnnotation(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
