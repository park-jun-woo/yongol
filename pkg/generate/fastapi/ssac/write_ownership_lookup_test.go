//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestWriteOwnershipLookup — writeOwnershipLookup owner row SELECT lookup 코드 출력 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteOwnershipLookup(t *testing.T) {
	var b strings.Builder
	ow := &ir.OwnershipInfo{
		Table:       "projects",
		OwnerColumn: "owner_id",
		ResourcePK:  "id",
	}
	writeOwnershipLookup(&b, ow, "    ")
	got := b.String()

	model := pascalCase(ir.DDLTableSingularIR("projects"))
	wants := []string{
		"    owner_row = await session.execute(",
		"select(" + model + ".owner_id).where(" + model + ".id == id)",
		"    owner = owner_row.scalar_one_or_none()",
	}
	for _, w := range wants {
		if !strings.Contains(got, w) {
			t.Errorf("expected %q in output:\n%s", w, got)
		}
	}
}
