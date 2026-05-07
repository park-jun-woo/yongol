// Const-only file (test data table).

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// pgtypeMatrixCases covers UUID / NUMERIC / TIMESTAMPTZ / TIMESTAMP /
// DATE / INET / CIDR / INTERVAL — every NeedsOverride=true family.
var pgtypeMatrixCases = []matrixCase{
	{"UUID NOT NULL", ddl.Column{RawType: "UUID", NotNull: true},
		"pgtype.UUID", "openapi_types.UUID", KindPgtype, true, true, "pgtypex.FromPgUUID"},
	{"UUID NULL", ddl.Column{RawType: "UUID"},
		"pgtype.UUID", "*openapi_types.UUID", KindPgtype, true, true, "pgtypex.FromPgUUIDPtr"},
	{"NUMERIC(10,2) NOT NULL", ddl.Column{RawType: "NUMERIC(10,2)", NotNull: true},
		"pgtype.Numeric", "string", KindPgtype, true, true, "pgtypex.FromPgNumeric"},
	{"DECIMAL(8,4) NULL", ddl.Column{RawType: "DECIMAL(8,4)"},
		"pgtype.Numeric", "*string", KindPgtype, true, true, "pgtypex.FromPgNumericPtr"},
	{"TIMESTAMPTZ NOT NULL", ddl.Column{RawType: "TIMESTAMPTZ", NotNull: true},
		"pgtype.Timestamptz", "time.Time", KindPgtype, true, true, "pgtypex.FromPgTimestamptz"},
	{"TIMESTAMP NOT NULL", ddl.Column{RawType: "TIMESTAMP", NotNull: true},
		"pgtype.Timestamp", "time.Time", KindPgtype, true, true, "pgtypex.FromPgTimestamp"},
	{"DATE NULL", ddl.Column{RawType: "DATE"},
		"pgtype.Date", "*time.Time", KindPgtype, true, true, "pgtypex.FromPgDatePtr"},
	{"INET NOT NULL", ddl.Column{RawType: "INET", NotNull: true},
		"pgtype.Inet", "string", KindPgtype, true, true, "pgInetToString"},
	{"CIDR NULL", ddl.Column{RawType: "CIDR"},
		"pgtype.Inet", "*string", KindPgtype, true, true, "pgInetToString"},
	{"INTERVAL NOT NULL", ddl.Column{RawType: "INTERVAL", NotNull: true},
		"pgtype.Interval", "string", KindPgtype, true, true, "pgtypex.FromPgInterval"},
}
