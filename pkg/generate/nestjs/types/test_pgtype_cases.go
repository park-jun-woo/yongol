// Const-only file (test data table) — filefunc skips //ff annotations
// on var/const-only files.

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// pgtypeBindCases covers the UUID / Numeric / TimestampTZ / Timestamp /
// Date / Inet / Interval / JSONB / Bytea families across NOT NULL and
// NULLABLE.
var pgtypeBindCases = []bindCase{
	// UUID
	{
		name:   "UUID_NotNull",
		family: typemap.FamilyUUID,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "String @db.Uuid", wantAPIType: "string",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "UUID_Nullable",
		family: typemap.FamilyUUID,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "String @db.Uuid", wantAPIType: "string | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Numeric
	{
		name:   "Numeric_NotNull",
		family: typemap.FamilyNumeric,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "Decimal", wantAPIType: "string",
		wantToDBExpr: "new Decimal({var})", wantToAPIExpr: "{row}.{field}.toString()",
		wantToRespExpr: "{var}.toString()", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "Numeric_Nullable",
		family: typemap.FamilyNumeric,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "Decimal", wantAPIType: "string | null",
		wantToDBExpr: "new Decimal({var})", wantToAPIExpr: "{row}.{field}.toString()",
		wantToRespExpr: "{var}.toString()", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// TimestampTZ
	{
		name:   "TimestampTZ_NotNull",
		family: typemap.FamilyTimestampTZ,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "DateTime", wantAPIType: "string",
		wantToDBExpr: "new Date({var})", wantToAPIExpr: "{row}.{field}.toISOString()",
		wantToRespExpr: "{var}.toISOString()", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "TimestampTZ_Nullable",
		family: typemap.FamilyTimestampTZ,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "DateTime", wantAPIType: "string | null",
		wantToDBExpr: "new Date({var})", wantToAPIExpr: "{row}.{field}.toISOString()",
		wantToRespExpr: "{var}.toISOString()", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Timestamp
	{
		name:   "Timestamp_NotNull",
		family: typemap.FamilyTimestamp,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "DateTime", wantAPIType: "string",
		wantToDBExpr: "new Date({var})", wantToAPIExpr: "{row}.{field}.toISOString()",
		wantToRespExpr: "{var}.toISOString()", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "Timestamp_Nullable",
		family: typemap.FamilyTimestamp,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "DateTime", wantAPIType: "string | null",
		wantToDBExpr: "new Date({var})", wantToAPIExpr: "{row}.{field}.toISOString()",
		wantToRespExpr: "{var}.toISOString()", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Date
	{
		name:   "Date_NotNull",
		family: typemap.FamilyDate,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "DateTime @db.Date", wantAPIType: "string",
		wantToDBExpr: "new Date({var})", wantToAPIExpr: "{row}.{field}.toISOString().slice(0,10)",
		wantToRespExpr: "{var}.toISOString().slice(0,10)", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "Date_Nullable",
		family: typemap.FamilyDate,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "DateTime @db.Date", wantAPIType: "string | null",
		wantToDBExpr: "new Date({var})", wantToAPIExpr: "{row}.{field}.toISOString().slice(0,10)",
		wantToRespExpr: "{var}.toISOString().slice(0,10)", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Inet
	{
		name:   "Inet_NotNull",
		family: typemap.FamilyInet,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "String", wantAPIType: "string",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "Inet_Nullable",
		family: typemap.FamilyInet,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "String", wantAPIType: "string | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Interval
	{
		name:   "Interval_NotNull",
		family: typemap.FamilyInterval,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "String", wantAPIType: "string",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "Interval_Nullable",
		family: typemap.FamilyInterval,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "String", wantAPIType: "string | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// JSONB
	{
		name:   "JSONB_NotNull",
		family: typemap.FamilyJSONB,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "Json", wantAPIType: "Record<string, unknown>",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "JSONB_Nullable",
		family: typemap.FamilyJSONB,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "Json", wantAPIType: "Record<string, unknown> | null",
		wantToDBExpr: "{var}", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
	// Bytea
	{
		name:   "Bytea_NotNull",
		family: typemap.FamilyBytea,
		opts:   ir.BindOpts{NotNull: true},
		wantDBType: "Bytes", wantAPIType: "Buffer",
		wantToDBExpr: "Buffer.from({var})", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "",
		wantSupported: true,
	},
	{
		name:   "Bytea_Nullable",
		family: typemap.FamilyBytea,
		opts:   ir.BindOpts{NotNull: false},
		wantDBType: "Bytes", wantAPIType: "Buffer | null",
		wantToDBExpr: "Buffer.from({var})", wantToAPIExpr: "{row}.{field}",
		wantToRespExpr: "{var}", wantNilCheck: "{var} === null",
		wantSupported: true,
	},
}
