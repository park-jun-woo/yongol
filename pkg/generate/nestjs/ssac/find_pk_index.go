//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what findPKIndex — FieldArg 배열에서 PK 인덱스 탐색

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// findPKIndex returns the index of the first arg whose key looks like a PK.
func findPKIndex(args []ir.FieldArg) int {
	for i, a := range args {
		key := resolveArgKey(a)
		if key == "id" || strings.HasSuffix(key, "_id") {
			return i
		}
	}
	return 0
}
