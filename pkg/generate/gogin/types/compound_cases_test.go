// Const-only file (test data table).

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// compoundMatrixCases covers JSONB / BYTEA / Array / Enum families.
var compoundMatrixCases = []matrixCase{
	{"JSONB NOT NULL", ddl.Column{RawType: "JSONB", NotNull: true},
		"[]byte", "map[string]interface{}", KindJSONB, false, true, ""},
	{"JSONB NULL", ddl.Column{RawType: "JSONB"},
		"[]byte", "*map[string]interface{}", KindJSONB, false, true, ""},
	{"JSON NOT NULL", ddl.Column{RawType: "JSON", NotNull: true},
		"[]byte", "map[string]interface{}", KindJSONB, false, true, ""},
	{"BYTEA NOT NULL", ddl.Column{RawType: "BYTEA", NotNull: true},
		"[]byte", "[]byte", KindBytea, false, true, ""},
	{"BYTEA NULL", ddl.Column{RawType: "BYTEA"},
		"[]byte", "[]byte", KindBytea, false, true, ""},
	{"TEXT[]", ddl.Column{RawType: "TEXT[]"},
		"[]string", "[]string", KindArray, false, true, ""},
	{"BIGINT[]", ddl.Column{RawType: "BIGINT[]"},
		"[]int64", "[]int64", KindArray, false, true, ""},
	{"BOOLEAN[]", ddl.Column{RawType: "BOOLEAN[]"},
		"[]bool", "[]bool", KindArray, false, true, ""},
	{"role enum NOT NULL", ddl.Column{
		RawType: "VARCHAR(20)", NotNull: true, VarcharLen: 20,
		CheckEnum: []string{"member", "admin"}},
		"string", "string", KindEnum, false, true, ""},
	{"role enum NULL", ddl.Column{
		RawType: "VARCHAR(20)", VarcharLen: 20,
		CheckEnum: []string{"member", "admin"}},
		"*string", "*string", KindEnum, false, true, ""},
}
