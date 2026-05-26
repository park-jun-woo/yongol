// Const-only file (test data table) — filefunc skips //ff annotations
// on var/const-only files.

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// compoundBindCases covers the Enum / Array / Unsupported families.
var compoundBindCases = []bindCase{
	// Enum
	{
		name:   "Enum_NotNull",
		family: typemap.FamilyEnum,
		opts:   ir.BindOpts{NotNull: true, EnumValues: []string{"active", "inactive"}},
		wantDBType: "String", wantAPIType: "string",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "Enum_Nullable",
		family: typemap.FamilyEnum,
		opts:   ir.BindOpts{NotNull: false, EnumValues: []string{"draft", "published"}},
		wantDBType: "String", wantAPIType: "string | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Array — TEXT[]
	{
		name:   "Array_Text_NotNull",
		family: typemap.FamilyArray,
		opts:   ir.BindOpts{NotNull: true, IsArray: true, ElementHead: "TEXT"},
		wantDBType: "String[]", wantAPIType: "string[]",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "Array_Text_Nullable",
		family: typemap.FamilyArray,
		opts:   ir.BindOpts{NotNull: false, IsArray: true, ElementHead: "TEXT"},
		wantDBType: "String[]", wantAPIType: "string[] | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Array — BIGINT[]
	{
		name:   "Array_Bigint_NotNull",
		family: typemap.FamilyArray,
		opts:   ir.BindOpts{NotNull: true, IsArray: true, ElementHead: "BIGINT"},
		wantDBType: "Int[]", wantAPIType: "number[]",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	// Array — BOOLEAN[]
	{
		name:   "Array_Boolean_NotNull",
		family: typemap.FamilyArray,
		opts:   ir.BindOpts{NotNull: true, IsArray: true, ElementHead: "BOOLEAN"},
		wantDBType: "Boolean[]", wantAPIType: "boolean[]",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	// Array — FLOAT8[]
	{
		name:   "Array_Float_NotNull",
		family: typemap.FamilyArray,
		opts:   ir.BindOpts{NotNull: true, IsArray: true, ElementHead: "FLOAT8"},
		wantDBType: "Float[]", wantAPIType: "number[]",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	// Array — unsupported element (UUID[])
	{
		name:   "Array_UUID_Unsupported",
		family: typemap.FamilyArray,
		opts:   ir.BindOpts{NotNull: true, IsArray: true, ElementHead: "UUID"},
		wantDBType: "/* unsupported array element: UUID */",
		wantAPIType: "/* unsupported array element: UUID */",
		wantToDBExpr: "", wantToAPIExpr: "",
		wantToRespExpr: "", wantNilCheck: "",
		wantSupported: false,
	},
	// Unsupported
	{
		name:   "Unsupported",
		family: typemap.FamilyUnsupported,
		opts:   ir.BindOpts{},
		wantDBType: "", wantAPIType: "",
		wantToDBExpr: "", wantToAPIExpr: "",
		wantToRespExpr: "", wantNilCheck: "",
		wantSupported: false,
	},
}
