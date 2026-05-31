//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what splitPutParts — PutOp Args 를 IsPK 기준 where/data Prisma 절 조각으로 분리

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// splitPutParts renders each arg as a "key: value" pair and partitions them
// into where parts (IsPK == true) and data parts.
func splitPutParts(args []ir.FieldArg) (whereParts, dataParts []string) {
	for _, a := range args {
		pair := fmt.Sprintf("%s: %s", resolveArgKey(a), renderArgValue(a))
		if a.IsPK {
			whereParts = append(whereParts, pair)
		} else {
			dataParts = append(dataParts, pair)
		}
	}
	return whereParts, dataParts
}
