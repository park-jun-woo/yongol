// Const-only file (test data table) — filefunc skips //ff annotations
// on var/const-only files.

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// nativeMatrixCases covers the integer / float / string / boolean
// families across NOT NULL, NULLABLE, NOT NULL DEFAULT.
var nativeMatrixCases = []matrixCase{
	{"BIGINT NOT NULL", ddl.Column{RawType: "BIGINT", NotNull: true},
		"int64", "int64", KindNative, false, true, "{row}.{field}"},
	{"INT NOT NULL", ddl.Column{RawType: "INT", NotNull: true},
		"int64", "int64", KindNative, false, true, ""},
	{"SMALLINT NOT NULL", ddl.Column{RawType: "SMALLINT", NotNull: true},
		"int64", "int64", KindNative, false, true, ""},
	{"BIGINT NULL", ddl.Column{RawType: "BIGINT"},
		"pgtype.Int8", "*int64", KindPgtype, false, true, "pgtypex.FromPgInt8Ptr"},
	{"BIGINT @nullable", ddl.Column{RawType: "BIGINT", NotNull: true, NullableAnnot: true},
		"pgtype.Int8", "*int64", KindPgtype, false, true, "pgtypex.FromPgInt8Ptr"},
	{"BIGINT NOT NULL DEFAULT 0", ddl.Column{RawType: "BIGINT", NotNull: true, HasDefault: true, DefaultLiteral: "0"},
		"int64", "int64", KindNative, false, true, ""},
	{"REAL NOT NULL", ddl.Column{RawType: "REAL", NotNull: true},
		"float64", "float64", KindNative, false, true, ""},
	{"FLOAT4 NULL", ddl.Column{RawType: "FLOAT4"},
		"pgtype.Float4", "*float64", KindPgtype, false, true, "pgtypex.FromPgFloat4Ptr"},
	{"FLOAT8 NULL", ddl.Column{RawType: "FLOAT8"},
		"pgtype.Float8", "*float64", KindPgtype, false, true, "pgtypex.FromPgFloat8Ptr"},
	{"VARCHAR(255) NOT NULL", ddl.Column{RawType: "VARCHAR(255)", NotNull: true, VarcharLen: 255},
		"string", "string", KindNative, false, true, ""},
	{"TEXT NOT NULL", ddl.Column{RawType: "TEXT", NotNull: true},
		"string", "string", KindNative, false, true, ""},
	{"TEXT NULL", ddl.Column{RawType: "TEXT"},
		"pgtype.Text", "*string", KindPgtype, false, true, "pgtypex.FromPgTextPtr"},
	{"CHAR(8) NOT NULL", ddl.Column{RawType: "CHAR(8)", NotNull: true, VarcharLen: 8},
		"string", "string", KindNative, false, true, ""},
	{"BOOLEAN NOT NULL", ddl.Column{RawType: "BOOLEAN", NotNull: true},
		"bool", "bool", KindNative, false, true, ""},
	{"BOOL NULL", ddl.Column{RawType: "BOOL"},
		"pgtype.Bool", "*bool", KindPgtype, false, true, "pgtypex.FromPgBoolPtr"},
}
