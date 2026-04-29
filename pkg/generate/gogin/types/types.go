//ff:func feature=gen-gogin type=util control=sequence
//ff:what MapPGType — ddl.Column 을 GoTypeBinding 으로 디스패치 (17 family + Unsupported)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// MapPGType returns the binding for a parsed DDL column. This is the
// single source of truth shared between codegen (Row→Model convert /
// INSERT params / response emit) and validate (Q-12, per-type Q-NN
// family, D-11). Every PG type the project supports is covered here;
// unrecognised heads route to KindUnsupported so D-11 can reject them
// before generate runs.
//
// The dispatch is purely structural — no inspection of the surrounding
// DDL or sqlc.yaml. If a column needs an override but the user's
// sqlc.yaml is missing it, the per-type Q-NN rule fires; this function
// always returns the canonical binding regardless.
func MapPGType(col ddl.Column) GoTypeBinding {
	notNull := isEffectivelyNotNull(col)
	def := col.DefaultLiteral
	info := parseRawType(col.RawType)
	return dispatchBinding(col, info, notNull, def)
}
