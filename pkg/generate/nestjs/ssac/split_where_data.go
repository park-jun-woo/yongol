//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what splitWhereData — FieldArg 배열 → Prisma where/data 분리

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// splitWhereData splits FieldArgs into where and data parts. The first arg (or
// one whose key matches common PK names) goes to where; the rest go to data.
func splitWhereData(args []ir.FieldArg) (string, string) {
	if len(args) == 0 {
		return "id: params.id", "...body"
	}
	pkIdx := findPKIndex(args)
	var whereParts, dataParts []string
	for i, a := range args {
		key := resolveArgKey(a)
		val := renderArgValue(a)
		pair := fmt.Sprintf("%s: %s", key, val)
		if i == pkIdx {
			whereParts = append(whereParts, pair)
		} else {
			dataParts = append(dataParts, pair)
		}
	}
	whereStr := strings.Join(whereParts, ", ")
	dataStr := strings.Join(dataParts, ", ")
	if dataStr == "" {
		dataStr = "...body"
	}
	return whereStr, dataStr
}
