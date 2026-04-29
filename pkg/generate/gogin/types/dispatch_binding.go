//ff:func feature=gen-gogin type=util control=selection
//ff:what dispatchBinding — MapPGType 의 family 분기 (CheckEnum/multi-token/array/pgtype/native/unsupported)

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// dispatchBinding routes a parsed column to the appropriate family
// constructor in priority order:
//
//  1. CheckEnum present → enum binding
//  2. Array marker → arrayBinding
//  3. pgtype family → mapPgtypeFamily
//  4. Native family → mapNativeFamily
//  5. Otherwise → unsupported (multi-word PG type not in the alias
//     matrix, or a CREATE TYPE user-defined ENUM)
//
// Multi-word PostgreSQL type names are no longer rejected up-front —
// parseRawType normalises recognised forms ("DOUBLE PRECISION" →
// "FLOAT8", "TIMESTAMP WITH TIME ZONE" → "TIMESTAMPTZ") so they reach
// the family matrices keyed by the single-token alias. Forms that do
// not appear in pgHeadAliases fall through to the final
// unsupportedBinding, which keeps D-11 firing on truly unknown heads.
//
// Extracted from MapPGType so each func stays inside the F1 line budget
// and the priority ladder is testable independently.
func dispatchBinding(col ddl.Column, info rawTypeInfo, notNull bool, def string) GoTypeBinding {
	switch {
	case len(col.CheckEnum) > 0:
		return enumBinding(notNull, def)
	case info.IsArray:
		return arrayBinding(info.Head, def)
	}
	if b, ok := mapPgtypeFamily(info.Head, notNull, def); ok {
		return b
	}
	if b, ok := mapNativeFamily(info.Head, notNull, def); ok {
		return b
	}
	return unsupportedBinding("PG type " + info.Head + " is not recognised (likely a CREATE TYPE user-defined ENUM, or a multi-word PG type without a registered single-token alias)")
}
