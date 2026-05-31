// Const-only file (test data table) — filefunc skips //ff annotations
// on var/const-only files.

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// nativeBindCases covers the Integer / Float / String / Boolean families
// across NOT NULL and NULLABLE.
var nativeBindCases = []bindCase{
	// Integer
	{
		name:       "Integer_NotNull",
		family:     typemap.FamilyInteger,
		opts:       ir.BindOpts{NotNull: true},
		wantDBType: "Int", wantAPIType: "number",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:       "Integer_Nullable",
		family:     typemap.FamilyInteger,
		opts:       ir.BindOpts{NotNull: false},
		wantDBType: "Int", wantAPIType: "number | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	{
		name:       "Integer_Default",
		family:     typemap.FamilyInteger,
		opts:       ir.BindOpts{NotNull: true, DefaultLiteral: "0"},
		wantDBType: "Int", wantAPIType: "number",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	// Float
	{
		name:       "Float_NotNull",
		family:     typemap.FamilyFloat,
		opts:       ir.BindOpts{NotNull: true},
		wantDBType: "Float", wantAPIType: "number",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:       "Float_Nullable",
		family:     typemap.FamilyFloat,
		opts:       ir.BindOpts{NotNull: false},
		wantDBType: "Float", wantAPIType: "number | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// String
	{
		name:       "String_NotNull",
		family:     typemap.FamilyString,
		opts:       ir.BindOpts{NotNull: true},
		wantDBType: "String", wantAPIType: "string",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:       "String_Nullable",
		family:     typemap.FamilyString,
		opts:       ir.BindOpts{NotNull: false},
		wantDBType: "String", wantAPIType: "string | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Boolean
	{
		name:       "Boolean_NotNull",
		family:     typemap.FamilyBoolean,
		opts:       ir.BindOpts{NotNull: true},
		wantDBType: "Boolean", wantAPIType: "boolean",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:       "Boolean_Nullable",
		family:     typemap.FamilyBoolean,
		opts:       ir.BindOpts{NotNull: false},
		wantDBType: "Boolean", wantAPIType: "boolean | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
}
