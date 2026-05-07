//ff:func feature=validate type=util control=selection topic=manifest-infra
//ff:what claimDDLTypeMatrix — XDN-06 claim 타입 ↔ DDL 타입 정합 매트릭스

package manifest_ddl

var claimDDLTypeMatrix = map[string][]string{
	"uuid":   {"UUID"},
	"string": {"TEXT", "VARCHAR", "CHARACTER VARYING", "CHAR", "BPCHAR"},
	"int64":  {"BIGINT", "INT8", "BIGSERIAL"},
	"int32":  {"INTEGER", "INT", "INT4", "SERIAL"},
	"bool":   {"BOOLEAN", "BOOL"},
}
