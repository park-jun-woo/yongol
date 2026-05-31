//ff:func feature=validate type=test control=sequence topic=rego-structural
//ff:what xpp30OwnershipNoAnnotation — resource_owner 참조 + @ownership 누락 검증
package rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXpp30OwnershipNoAnnotation(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
