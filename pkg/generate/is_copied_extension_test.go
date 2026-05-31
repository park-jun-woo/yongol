//ff:func feature=generate type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"testing"
)

func TestIsCopiedExtension(t *testing.T) {
	for _, p := range []string{"a.tsx", "b.TS", "c.jsx", "d.js", "e.css", "f.svg"} {
		if !isCopiedExtension(p) {
			t.Errorf("isCopiedExtension(%q) = false, want true", p)
		}
	}
	for _, p := range []string{"x.go", "y.json", "z"} {
		if isCopiedExtension(p) {
			t.Errorf("isCopiedExtension(%q) = true, want false", p)
		}
	}
}
