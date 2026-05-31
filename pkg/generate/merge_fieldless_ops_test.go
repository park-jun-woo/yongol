//ff:func feature=generate type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"testing"
)

func TestMergeFieldlessOps(t *testing.T) {
	dst := map[string]bool{"A": true}
	src := map[string]bool{"B": true, "C": true}
	mergeFieldlessOps(dst, src)
	for _, k := range []string{"A", "B", "C"} {
		if !dst[k] {
			t.Errorf("dst missing %q after merge", k)
		}
	}
}
