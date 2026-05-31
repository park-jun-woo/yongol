//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what splitPKArgs — FieldArg 목록을 IsPK 기준 where 인자와 data 인자로 분리

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// splitPKArgs partitions args into where args (IsPK == true) and data args.
func splitPKArgs(args []ir.FieldArg) (whereArgs, dataArgs []ir.FieldArg) {
	for _, a := range args {
		if a.IsPK {
			whereArgs = append(whereArgs, a)
		} else {
			dataArgs = append(dataArgs, a)
		}
	}
	return whereArgs, dataArgs
}
