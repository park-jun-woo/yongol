//ff:func feature=generate type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"testing"
)

func TestIsYongolManaged(t *testing.T) {
	managed := []string{
		"src/api.ts",
		"src/types/Course.ts",
		"src/lib/http.ts",
		"src/components/ui/button.tsx",
	}
	for _, p := range managed {
		if !isYongolManaged(p) {
			t.Errorf("isYongolManaged(%q) = false, want true", p)
		}
	}
	unmanaged := []string{
		"src/pages/Home.tsx",
		"src/components/MyCard.tsx",
		"README.md",
	}
	for _, p := range unmanaged {
		if isYongolManaged(p) {
			t.Errorf("isYongolManaged(%q) = true, want false", p)
		}
	}
}
