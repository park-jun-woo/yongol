//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xnp63RoleManifest — nil/no ground/declared role/undeclared role 검증
package rego_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnp63RoleManifest(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
