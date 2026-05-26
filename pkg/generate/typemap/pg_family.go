//ff:type feature=gen-typemap type=model
//ff:what PGFamily — PostgreSQL 타입 family 열거형 (14 family + Array + Unsupported)

package typemap

// PGFamily classifies a PostgreSQL column type into one of the canonical
// families used by code generators across all backend targets. The
// classification is language-agnostic: the same PGFamily applies whether
// the target is Go+Gin, NestJS, or FastAPI.
type PGFamily int

const (
	// native families (mapNativeFamily)

	// FamilyInteger groups BIGINT, INT, SMALLINT, SERIAL and all their
	// aliases. All map to a single integer representation (e.g. int64 in Go).
	FamilyInteger PGFamily = iota
	// FamilyFloat groups REAL, FLOAT, FLOAT4, FLOAT8.
	FamilyFloat
	// FamilyString groups VARCHAR, TEXT, CHAR, BPCHAR.
	FamilyString
	// FamilyBoolean groups BOOLEAN, BOOL.
	FamilyBoolean

	// pgtype families (mapPgtypeFamily)

	// FamilyUUID represents the UUID type.
	FamilyUUID
	// FamilyNumeric groups NUMERIC, DECIMAL.
	FamilyNumeric
	// FamilyTimestamp represents TIMESTAMP (without time zone).
	FamilyTimestamp
	// FamilyTimestampTZ represents TIMESTAMPTZ (with time zone).
	FamilyTimestampTZ
	// FamilyDate represents the DATE type.
	FamilyDate
	// FamilyInet groups INET, CIDR.
	FamilyInet
	// FamilyInterval represents the INTERVAL type.
	FamilyInterval
	// FamilyJSONB groups JSONB, JSON.
	FamilyJSONB
	// FamilyBytea represents the BYTEA type.
	FamilyBytea

	// special dispatch paths

	// FamilyEnum is selected when CheckEnum is non-empty.
	FamilyEnum
	// FamilyArray is selected when the column type ends with [].
	FamilyArray
	// FamilyUnsupported is the fallback for unrecognised types.
	FamilyUnsupported
)
