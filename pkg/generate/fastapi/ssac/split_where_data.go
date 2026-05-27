//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what splitWhereData — FieldArg 배열 → SQLAlchemy where/data 분리

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// splitWhereData splits FieldArgs into where and data parts. The arg whose
// key matches a common PK pattern goes to where; the rest go to data.
func splitWhereData(args []ir.FieldArg) (where []ir.FieldArg, data []ir.FieldArg) {
	if len(args) == 0 {
		return nil, nil
	}
	pkIdx := findPKIndex(args)
	for i, a := range args {
		if i == pkIdx {
			where = append(where, a)
		} else {
			data = append(data, a)
		}
	}
	return where, data
}
