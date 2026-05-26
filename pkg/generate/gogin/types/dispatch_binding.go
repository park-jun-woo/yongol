//ff:func feature=gen-gogin type=util control=selection
//ff:what dispatchBinding — MapPGType 의 family 분기 (typemap.ClassifyFamily 위임 → Go 바인딩 산출)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// dispatchBinding routes a parsed column to the appropriate family
// constructor by delegating family classification to
// typemap.ClassifyFamily and then mapping the resulting PGFamily to the
// Go-specific binding.
//
// Multi-word PostgreSQL type names are normalised by typemap.ParseRawType
// ("DOUBLE PRECISION" → "FLOAT8", "TIMESTAMP WITH TIME ZONE" →
// "TIMESTAMPTZ") so the downstream binding functions receive the
// single-token alias.
//
// Extracted from MapPGType so each func stays inside the F1 line budget
// and the priority ladder is testable independently.
func dispatchBinding(col ddl.Column, info rawTypeInfo, notNull bool, def string) GoTypeBinding {
	family := typemap.ClassifyFamily(columnAdapter{col})
	switch family {
	case typemap.FamilyEnum:
		return enumBinding(notNull, def)
	case typemap.FamilyArray:
		return arrayBinding(info.Head, def)
	case typemap.FamilyUUID:
		return pgtypeUUID(notNull, def)
	case typemap.FamilyNumeric:
		return pgtypeNumeric(notNull, def)
	case typemap.FamilyTimestampTZ:
		return pgtypeTimestamp("TIMESTAMPTZ", notNull, def)
	case typemap.FamilyTimestamp:
		return pgtypeTimestamp("TIMESTAMP", notNull, def)
	case typemap.FamilyDate:
		return pgtypeTimestamp("DATE", notNull, def)
	case typemap.FamilyInet:
		return pgtypeInet(notNull, def)
	case typemap.FamilyInterval:
		return pgtypeInterval(notNull, def)
	case typemap.FamilyJSONB:
		return jsonbBinding(notNull, def)
	case typemap.FamilyBytea:
		return byteaBinding(notNull, def)
	case typemap.FamilyInteger:
		return nativeInteger(notNull, def)
	case typemap.FamilyFloat:
		return nativeFloatWithHead(info.Head, notNull, def)
	case typemap.FamilyString:
		return nativeString(notNull, def)
	case typemap.FamilyBoolean:
		return nativeBoolean(notNull, def)
	default:
		return unsupportedBinding("PG type " + info.Head + " is not recognised (likely a CREATE TYPE user-defined ENUM, or a multi-word PG type without a registered single-token alias)")
	}
}
