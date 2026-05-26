//ff:func feature=gen-typemap type=util control=selection
//ff:what classifyPgtypeFamily — pgtype 기반 head 토큰을 PGFamily 로 매핑

package typemap

// classifyPgtypeFamily maps pgtype-backed head tokens to their PGFamily.
func classifyPgtypeFamily(head string) (PGFamily, bool) {
	switch head {
	case "UUID":
		return FamilyUUID, true
	case "NUMERIC", "DECIMAL":
		return FamilyNumeric, true
	case "TIMESTAMPTZ":
		return FamilyTimestampTZ, true
	case "TIMESTAMP":
		return FamilyTimestamp, true
	case "DATE":
		return FamilyDate, true
	case "INET", "CIDR":
		return FamilyInet, true
	case "INTERVAL":
		return FamilyInterval, true
	case "JSONB", "JSON":
		return FamilyJSONB, true
	case "BYTEA":
		return FamilyBytea, true
	}
	return FamilyUnsupported, false
}
