//ff:func feature=gen-fastapi type=util control=selection
//ff:what bindFastAPI — PGFamily + BindOpts → ir.TypeBinding 디스패치 (family 별 바인딩 위임)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindFastAPI dispatches to the per-family binding functions based on the
// PGFamily classification. Each helper returns an ir.TypeBinding with
// SQLAlchemy column type (DBType), Python/Pydantic API type (APIType),
// and the conversion expressions between DB and API layers.
func bindFastAPI(family typemap.PGFamily, opts ir.BindOpts) ir.TypeBinding {
	switch family {
	case typemap.FamilyInteger:
		return bindInteger(opts)
	case typemap.FamilyFloat:
		return bindFloat(opts)
	case typemap.FamilyString:
		return bindString(opts)
	case typemap.FamilyBoolean:
		return bindBoolean(opts)
	case typemap.FamilyUUID:
		return bindUUID(opts)
	case typemap.FamilyNumeric:
		return bindNumeric(opts)
	case typemap.FamilyTimestampTZ:
		return bindTimestampTZ(opts)
	case typemap.FamilyTimestamp:
		return bindTimestamp(opts)
	case typemap.FamilyDate:
		return bindDate(opts)
	case typemap.FamilyInet:
		return bindInet(opts)
	case typemap.FamilyInterval:
		return bindInterval(opts)
	case typemap.FamilyJSONB:
		return bindJSONB(opts)
	case typemap.FamilyBytea:
		return bindBytea(opts)
	case typemap.FamilyEnum:
		return bindEnum(opts)
	case typemap.FamilyArray:
		return bindArray(opts)
	default:
		return bindUnsupported(family)
	}
}
