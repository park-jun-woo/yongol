//ff:func feature=gen-gogin type=util control=selection
//ff:what dispatchBinding — MapPGType 의 family 분기 (CheckEnum/multi-token/array/pgtype/native/unsupported)

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// dispatchBinding routes a parsed column to the appropriate family
// constructor in priority order:
//
//  1. CheckEnum present → enum binding
//  2. Multi-token raw type → unsupported (parser limit)
//  3. Array marker → arrayBinding
//  4. pgtype family → mapPgtypeFamily
//  5. Native family → mapNativeFamily
//  6. Otherwise → unsupported (likely CREATE TYPE user-defined ENUM)
//
// Extracted from MapPGType so each func stays inside the F1 line budget
// and the priority ladder is testable independently.
func dispatchBinding(col ddl.Column, info rawTypeInfo, notNull bool, def string) GoTypeBinding {
	switch {
	case len(col.CheckEnum) > 0:
		return enumBinding(notNull, def)
	case info.MultiToken:
		return unsupportedBinding("multi-word PG type " + info.Head + " is not supported (use the single-token alias such as TIMESTAMPTZ instead of TIMESTAMP WITH TIME ZONE)")
	case info.IsArray:
		return arrayBinding(info.Head, def)
	}
	if b, ok := mapPgtypeFamily(info.Head, notNull, def); ok {
		return b
	}
	if b, ok := mapNativeFamily(info.Head, notNull, def); ok {
		return b
	}
	return unsupportedBinding("PG type " + info.Head + " is not recognised (likely a CREATE TYPE user-defined ENUM — not yet supported)")
}
