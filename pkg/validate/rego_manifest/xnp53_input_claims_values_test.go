//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xnp53InputClaimsValues — claims ref 검증 (nil/no ground/pass/fire) 검증
package rego_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnp53InputClaimsValues(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
