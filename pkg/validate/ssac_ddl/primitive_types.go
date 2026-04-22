//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what primitiveTypes — DDL 테이블 매칭에서 제외할 Go 기본 타입 집합

package ssac_ddl

// primitiveTypes are Go types that never map to DDL tables.
var primitiveTypes = map[string]bool{
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
	"float32": true, "float64": true,
	"string": true, "bool": true, "byte": true, "rune": true,
	"error": true, "any": true,
}
