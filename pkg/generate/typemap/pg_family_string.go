//ff:func feature=gen-typemap type=util control=selection
//ff:what String — PGFamily 의 사람이 읽기 편한 문자열 표현

package typemap

// String returns a human-readable label for the PGFamily value.
func (f PGFamily) String() string {
	switch f {
	case FamilyInteger:
		return "Integer"
	case FamilyFloat:
		return "Float"
	case FamilyString:
		return "String"
	case FamilyBoolean:
		return "Boolean"
	case FamilyUUID:
		return "UUID"
	case FamilyNumeric:
		return "Numeric"
	case FamilyTimestamp:
		return "Timestamp"
	case FamilyTimestampTZ:
		return "TimestampTZ"
	case FamilyDate:
		return "Date"
	case FamilyInet:
		return "Inet"
	case FamilyInterval:
		return "Interval"
	case FamilyJSONB:
		return "JSONB"
	case FamilyBytea:
		return "Bytea"
	case FamilyEnum:
		return "Enum"
	case FamilyArray:
		return "Array"
	case FamilyUnsupported:
		return "Unsupported"
	default:
		return "Unknown"
	}
}
