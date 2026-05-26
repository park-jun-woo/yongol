//ff:func feature=gen-typemap type=util control=sequence
//ff:what ClassifyFamily — ColumnMeta 를 PGFamily 로 분류 (CheckEnum/Array/pgtype/native/Unsupported)

package typemap

// ClassifyFamily determines the PGFamily for a DDL column using the same
// priority order as the gogin/types dispatcher:
//
//  1. CheckEnum non-empty → FamilyEnum
//  2. Array marker → FamilyArray
//  3. pgtype family (UUID, NUMERIC, TIMESTAMP, INET, INTERVAL, JSONB, BYTEA)
//  4. Native family (Integer, Float, String, Boolean)
//  5. Otherwise → FamilyUnsupported
func ClassifyFamily(col ColumnMeta) PGFamily {
	info := ParseRawType(col.RawType())

	if len(col.CheckEnum()) > 0 {
		return FamilyEnum
	}
	if info.IsArray {
		return FamilyArray
	}
	if f, ok := classifyPgtypeFamily(info.Head); ok {
		return f
	}
	if f, ok := classifyNativeFamily(info.Head); ok {
		return f
	}
	return FamilyUnsupported
}
