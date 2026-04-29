// Const-only file (test data table).

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// unsupportedMatrixCases covers types the dispatcher must reject:
// multi-word PG types whose single-token alias is not yet present in
// any family matrix (TIME / TIMETZ / VARBIT) and unrecognised heads
// (likely CREATE TYPE user-defined ENUMs). DOUBLE PRECISION and
// TIMESTAMP WITH TIME ZONE used to live here but are now normalised
// to FLOAT8 / TIMESTAMPTZ via ddl.NormalizePGTypeHead and dispatch
// successfully.
var unsupportedMatrixCases = []matrixCase{
	{"TIME WITH TIME ZONE", ddl.Column{RawType: "TIME WITH TIME ZONE", NotNull: true},
		"", "", KindUnsupported, false, false, ""},
	{"TIME WITHOUT TIME ZONE", ddl.Column{RawType: "TIME WITHOUT TIME ZONE", NotNull: true},
		"", "", KindUnsupported, false, false, ""},
	{"BIT VARYING", ddl.Column{RawType: "BIT VARYING(8)", NotNull: true},
		"", "", KindUnsupported, false, false, ""},
	{"USER_DEFINED_ENUM", ddl.Column{RawType: "ORDER_STATUS", NotNull: true},
		"", "", KindUnsupported, false, false, ""},
}
