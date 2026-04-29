//ff:func feature=gen-gogin type=util control=selection
//ff:what isEffectivelyNotNull — NOT NULL constraint 와 @nullable annotation 결합

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// isEffectivelyNotNull combines the explicit NOT NULL constraint with
// the user-visible @nullable annotation. The annotation overrides the
// constraint for downstream Go-type decisions; an annotated column is
// always treated as nullable even if the DDL forgot the modifier.
func isEffectivelyNotNull(col ddl.Column) bool {
	switch {
	case col.NullableAnnot:
		return false
	default:
		return col.NotNull
	}
}
