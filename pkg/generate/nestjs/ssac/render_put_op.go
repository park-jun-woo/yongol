//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderPutOp — PutOp → IsPK 기반 where/data 분리 Prisma update 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPutOp writes a Prisma update call. Args with IsPK == true go to the
// where clause; the rest go to the data clause. This uses the Phase018 IR
// enrichment instead of heuristic PK detection.
func renderPutOp(b *strings.Builder, op *ir.PutOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)

	var whereParts, dataParts []string
	for _, a := range op.Args {
		key := resolveArgKey(a)
		val := renderArgValue(a)
		pair := fmt.Sprintf("%s: %s", key, val)
		if a.IsPK {
			whereParts = append(whereParts, pair)
		} else {
			dataParts = append(dataParts, pair)
		}
	}

	whereStr := strings.Join(whereParts, ", ")
	if whereStr == "" {
		whereStr = "id: params.id"
	}
	dataStr := strings.Join(dataParts, ", ")
	if dataStr == "" {
		dataStr = "...body"
	}

	b.WriteString(fmt.Sprintf("%sawait %s.%s.update({ where: { %s }, data: { %s } });\n",
		indent, prismaRef, model, whereStr, dataStr))
}
