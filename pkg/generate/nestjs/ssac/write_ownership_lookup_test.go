//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestWriteOwnershipLookup — Ownership 메타 기반 owner findUnique lookup 코드 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteOwnershipLookup(t *testing.T) {
	var b strings.Builder
	ow := &ir.OwnershipInfo{
		Table:       "notes",
		OwnerColumn: "owner_id",
		ResourcePK:  "id",
	}
	writeOwnershipLookup(&b, ow, "  ", "this.prisma")

	out := b.String()
	for _, want := range []string{
		"const owner = await this.prisma.note.findUnique({\n",
		"  where: { id: params.id },\n",
		"  select: { owner_id: true },\n",
		"});\n",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("lookup missing %q\n--- got ---\n%s", want, out)
		}
	}
}
