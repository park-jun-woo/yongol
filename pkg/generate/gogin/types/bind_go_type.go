//ff:func feature=gen-gogin type=util control=selection
//ff:what bindGoType — PGFamily + BindOpts → GoTypeBinding 디스패치 (기존 family 생성자 위임)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindGoType dispatches to the existing family constructors using opts.
func bindGoType(family typemap.PGFamily, opts ir.BindOpts) GoTypeBinding {
	notNull := opts.NotNull
	def := opts.DefaultLiteral

	switch family {
	case typemap.FamilyEnum:
		return enumBinding(notNull, def)
	case typemap.FamilyArray:
		return arrayBinding(opts.ElementHead, def)
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
		return nativeFloatWithHead(opts.ElementHead, notNull, def)
	case typemap.FamilyString:
		return nativeString(notNull, def)
	case typemap.FamilyBoolean:
		return nativeBoolean(notNull, def)
	default:
		return unsupportedBinding("unrecognised PG family")
	}
}
