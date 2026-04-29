// Const-only file (test data table).

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// unsupportedMatrixCases covers types the dispatcher must reject:
// multi-token PG types and unrecognised heads (likely CREATE TYPE
// user-defined ENUM).
var unsupportedMatrixCases = []matrixCase{
	{"DOUBLE PRECISION", ddl.Column{RawType: "DOUBLE PRECISION", NotNull: true},
		"", "", KindUnsupported, false, false, ""},
	{"TIMESTAMP WITH TIME ZONE", ddl.Column{RawType: "TIMESTAMP WITH TIME ZONE", NotNull: true},
		"", "", KindUnsupported, false, false, ""},
	{"USER_DEFINED_ENUM", ddl.Column{RawType: "ORDER_STATUS", NotNull: true},
		"", "", KindUnsupported, false, false, ""},
}
