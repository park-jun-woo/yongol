//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderGetOp — GetOp → Prisma findUnique/findMany + PaginationArgs 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderGetOp writes a Prisma findUnique or findMany call. PaginationArgs
// from the Phase018 IR are rendered as Prisma take/skip/cursor options
// separate from the where clause. Variable shadowing is already resolved
// in the IR (Phase018), so VarName is used directly.
func renderGetOp(b *strings.Builder, op *ir.GetOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)
	method := "findUnique"
	if op.IsList {
		method = "findMany"
	}

	// Build Prisma query options.
	var opts []string

	// where clause from Args.
	if len(op.Args) > 0 {
		whereParts := make([]string, 0, len(op.Args))
		for _, a := range op.Args {
			key := resolveArgKey(a)
			whereParts = append(whereParts, fmt.Sprintf("%s: %s", key, renderArgValue(a)))
		}
		opts = append(opts, fmt.Sprintf("where: { %s }", strings.Join(whereParts, ", ")))
	}

	// Pagination options from PaginationArgs (cursor, per_page, etc.).
	for _, pa := range op.PaginationArgs {
		key := resolveArgKey(pa)
		val := renderArgValue(pa)
		switch key {
		case "per_page", "limit":
			opts = append(opts, fmt.Sprintf("take: %s", val))
		case "page_offset", "offset":
			opts = append(opts, fmt.Sprintf("skip: %s", val))
		case "cursor":
			opts = append(opts, fmt.Sprintf("cursor: %s ? { id: %s } : undefined", val, val))
		default:
			opts = append(opts, fmt.Sprintf("%s: %s", key, val))
		}
	}

	var argsStr string
	if len(opts) > 0 {
		argsStr = "{ " + strings.Join(opts, ", ") + " }"
	} else {
		argsStr = "{}"
	}

	b.WriteString(fmt.Sprintf("%sconst %s = await %s.%s.%s(%s);\n",
		indent, op.VarName, prismaRef, model, method, argsStr))
}
