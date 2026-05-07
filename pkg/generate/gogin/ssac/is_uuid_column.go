//ff:func feature=gen-gogin type=util control=sequence
//ff:what isUUIDColumn — 컬럼이 UUID 바인딩인지 판별

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// isUUIDColumn returns true when the column resolves to a UUID binding.
func isUUIDColumn(col *ddl.Column) bool {
	binding := types.MapPGType(*col)
	return binding.SqlcGoType == "pgtype.UUID"
}
